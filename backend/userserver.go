package main

import (
	"context"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ahmetb/sorucevap/genpb/sorucevap"
)

const (
	usersCol = `users`
)

type MyUserService struct {
	fsClient *firestore.Client
}

func (m MyUserService) GetUser(ctx context.Context, request *sorucevap.GetUserRequest) (*sorucevap.UserRecord, error) {
	doc, err := m.fsClient.Collection(usersCol).Doc(request.GetId()).Get(ctx)
	if status.Code(err) == codes.NotFound {
		return nil, status.Errorf(codes.NotFound, "user not found")
	} else if err != nil {
		log.Printf("failed to get user: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get user")
	}
	user := &sorucevap.UserRecord{}
	MapToProto(doc.Data(), user)
	return user, nil
}

func (m MyUserService) AddUser(ctx context.Context, record *sorucevap.UserRecord) (*empty.Empty, error) {
	if record.GetID() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "id not specified")
	}

	record.RegistrationDate = time.Now().Unix()

	_, err := m.fsClient.Collection(usersCol).Doc(record.GetID()).Set(ctx, ProtoToMap(record))
	if err != nil {
		log.Printf("failed to get user: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to save user")
	}
	return &empty.Empty{}, nil
}
