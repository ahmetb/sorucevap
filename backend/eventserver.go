package main

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ahmetb/sorucevap/genpb/sorucevap"
)

const (
	eventsCol = "events"
)

type MyEventServer struct {
	fsClient *firestore.Client
}

func (m MyEventServer) GetEvent(ctx context.Context, request *sorucevap.GetEventRequest) (*sorucevap.Event, error) {
	doc, err := m.fsClient.Collection(eventsCol).Doc(request.GetId()).Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "event %s not found", request.GetId())
		}
		log.Printf("delete error: %v (%T)", err, err)
		return nil, status.Errorf(codes.Internal, "query error")
	}
	event := new(sorucevap.Event)
	MapToProto(doc.Data(), event)
	return event, nil
}

func (m MyEventServer) CreateEvent(ctx context.Context, event *sorucevap.Event) (*sorucevap.Event, error) {
	event.Id = uuid.New().String()
	eventV := ProtoToMap(event)
	_, err := m.fsClient.Collection(eventsCol).Doc(event.Id).Set(ctx, eventV)
	if err != nil {
		log.Printf("failed to insert event: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to insert event")
	}
	return event, nil
}

func (m MyEventServer) DeleteEvent(ctx context.Context, request *sorucevap.DeleteEventRequest) (*empty.Empty, error) {
	_, err := m.fsClient.Collection(eventsCol).Doc(request.GetId()).Delete(ctx, firestore.Exists)
	if err == nil {
		return &empty.Empty{}, nil
	}
	if status.Code(err) == codes.NotFound {
		return nil, status.Errorf(codes.NotFound, "event %s not found", request.GetId())
	}
	log.Printf("delete error: %v (%T)", err, err)
	return nil, status.Errorf(codes.Internal, "delete error")
}
