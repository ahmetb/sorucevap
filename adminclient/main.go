package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	"github.com/ahmetb/sorucevap/genpb/sorucevap"
)

func main() {
	ctx := context.Background()
	dialCtx, cleanup := context.WithTimeout(ctx, time.Second*15)
	defer cleanup()

	conn, err := grpc.DialContext(dialCtx, "localhost:8080",
		grpc.WithInsecure())
	defer conn.Close()

	if err != nil {
		log.Fatal(err)
	}
	client := sorucevap.NewEventsClient(conn)

	// create event
	ev, err := client.CreateEvent(ctx, &sorucevap.Event{
		Name:        "deneme",
		Description: "deneme aciklama",
		ExpiresAt:   time.Now().Add(time.Hour * 48).Unix(),
	})
	if err != nil {
		log.Fatalf("failed to create event: %v", err)
	}
	log.Printf("Created event: %s", ev.GetId())

	time.Sleep(time.Second*2)


	// get event
	ev, err = client.GetEvent(ctx, &sorucevap.GetEventRequest{
		Id: ev.GetId(),
	})
	if err != nil {
		log.Fatalf("failed to get event: %v", err)
	}
	log.Printf("event queried = %#v", ev)

	time.Sleep(time.Second*2)

	// delete event
	_, err = client.DeleteEvent(ctx, &sorucevap.DeleteEventRequest{
		Id: ev.GetId(),
	})
	if err != nil {
		log.Fatalf("failed to delete event#%s: %v", ev.GetId(), err)
	}
	log.Fatalf("deleted event %s", ev.GetId())
}
