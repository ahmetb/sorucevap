package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	_ "cloud.google.com/go/firestore"
	"github.com/fatih/color"
	"google.golang.org/grpc"

	"github.com/ahmetb/sorucevap/genpb/sorucevap"
)

func main() {
	ctx := context.Background()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := os.Getenv("LISTEN_ADDR")
	listenAddr := fmt.Sprintf("%s:%s", addr, port)
	lis, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("failed to listen on %s: %v", listenAddr, err)
	}

	fsClient, err := firestore.NewClient(ctx, "ahmetb-demo") // TODO detect
	if err != nil {
		log.Fatalf("failed to initialize datastore client: %v", err)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(withLogging))

	eventServer := MyEventServer{fsClient: fsClient}
	sorucevap.RegisterEventsServer(grpcServer, eventServer)
	userServer := MyUserService{fsClient: fsClient}
	sorucevap.RegisterUsersServer(grpcServer, userServer)

	log.Println("starting grpc server")
	log.Fatal(grpcServer.Serve(lis))
}

func withLogging(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (resp interface{}, err error) {

	start := time.Now()
	defer func() {
		color.New(color.FgHiBlack).Printf("request(%s) completed in %v\n", info.FullMethod, time.Since(start))
	}()
	return handler(ctx, req)
}
