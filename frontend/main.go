package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ahmetb/sorucevap/genpb/sorucevap"
)

var (
	oauth2Cfg = &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RedirectURL:  "http://localhost:8081/authorize",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := os.Getenv("LISTEN_ADDR")
	listenAddr := fmt.Sprintf("%s:%s", addr, port)

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/mylogin", loginRedirect)
	mux.HandleFunc("/login", loginHandler)
	mux.HandleFunc("/logout", logoutHandler)
	mux.HandleFunc("/authorize", authorizeHandler)
	log.Println("starting server")

	handler := authenticatingMiddleware(mux)
	http.ListenAndServe(listenAddr, handler)
}

func authenticatingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		c, err := req.Cookie("jwt")
		if req.URL.Path == "/logout" || req.URL.Path == "/login" || err != nil {
			next.ServeHTTP(w, req)
			return
		}

		t, err := verifyIDToken(req.Context(), c.Value)
		if err != nil {
			log.Println(err)
			http.Redirect(w, req, "/logout", http.StatusFound)
			return
		}
		log.Printf("path=%s - authenticated as user %q", req.URL.Path, t.Claims["name"].(string))
		ctx := context.WithValue(req.Context(), "token", t)
		req = req.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}

func loginRedirect(w http.ResponseWriter, req *http.Request) {
	url := oauth2Cfg.AuthCodeURL("state")
	http.Redirect(w, req, url, http.StatusTemporaryRedirect)
}

func authorizeHandler(w http.ResponseWriter, req *http.Request) {
	code := req.URL.Query().Get("code")
	tok, err := oauth2Cfg.Exchange(req.Context(), code)
	if err != nil {
		log.Fatal(err)
	}
	_ = tok.Extra("id_token")
}

func logoutHandler(w http.ResponseWriter, req *http.Request) {
	for _, c := range req.Cookies() {
		c.Expires = time.Now().Add(time.Hour * -24)
		http.SetCookie(w, c)
	}
	http.Redirect(w, req, "/", http.StatusFound)
}

func home(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" || req.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	f, err := os.Open("index.html")
	if err != nil {
		writeError(w, err)
		return
	}
	io.Copy(w, f)
}

func loginHandler(w http.ResponseWriter, req *http.Request) {
	idToken := req.URL.Query().Get("jwt")

	token, err := verifyIDToken(req.Context(), idToken)
	if err != nil {
		writeError(w, fmt.Errorf("failed to verify ID token: %w", err))
		return
	}
	user := &sorucevap.UserRecord{
		ID:                token.Claims["user_id"].(string),
		FullName:          token.Claims["name"].(string),
		ProfilePictureURL: token.Claims["picture"].(string),
	}

	cc, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		writeError(w, fmt.Errorf("grpc error: %w", err))
		return
	}
	cl := sorucevap.NewUsersClient(cc)
	_, err = cl.GetUser(req.Context(), &sorucevap.GetUserRequest{Id: user.GetID()})
	if err == nil {
		log.Println("user exists!")
		goto ret
	}

	if status.Code(err) != codes.NotFound {
		writeError(w, err)
		return
	}

	// user does not exist, create
	if _, err := cl.AddUser(req.Context(), user); err != nil {
		writeError(w, err)
		return
	}
	log.Println("registered user!")

ret:
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    idToken,
		Path:     "",
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: false,
		SameSite: http.SameSiteStrictMode,
	})
	http.Redirect(w, req, "/", http.StatusFound)
}

func verifyIDToken(ctx context.Context, jwtToken string) (*auth.Token, error) {
	app, err := firebase.NewApp(ctx, nil,
		option.WithCredentialsFile("/Users/ahmetb/Downloads/"+
			"ahmetb-demo-firebase-adminsdk-s64ld-2c3d6e0064.json"))
	if err != nil {
		return nil, fmt.Errorf("failed to initialize firebase app: %v", err)
	}

	client, err := app.Auth(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting Auth client: %v", err)
	}

	token, err := client.VerifyIDToken(ctx, jwtToken)
	if err != nil {
		return nil, fmt.Errorf("error verifying ID token: %v", err)
	}
	return token, nil
}

func writeError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, "error: %v", err)
}
