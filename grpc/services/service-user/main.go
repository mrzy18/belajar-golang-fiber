package main

import (
	"context"
	"log"
	"net"

	"grpc/common/config"
	"grpc/common/model"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

var localStorage *model.UserList

func init() {
	localStorage = new(model.UserList)
	localStorage.List = make([]*model.User, 0)
}

type UsersServer struct {
	model.UnimplementedUsersServer
}

func (UsersServer) Register(_ context.Context, user *model.User) (*emptypb.Empty, error) {
	localStorage.List = append(localStorage.List, user)

	log.Println("Registering user", user.String())

	return new(emptypb.Empty), nil
}

func (UsersServer) List(context.Context, *emptypb.Empty) (*model.UserList, error) {
	return localStorage, nil
}

func main() {
	srv := grpc.NewServer()
	var userSrv UsersServer
	model.RegisterUsersServer(srv, userSrv)

	log.Println("Starting RPC server at", config.ServiceUserPort)

	l, err := net.Listen("tcp", config.ServiceUserPort)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", config.ServiceUserPort, err)
	}

	log.Fatal(srv.Serve(l))
}
