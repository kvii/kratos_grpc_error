package main

import (
	"context"
	"log"

	"github.com/go-kratos/examples/errors/api"
	"github.com/go-kratos/examples/helloworld/helloworld"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

type Service struct {
	helloworld.UnimplementedGreeterServer
}

func (Service) SayHello(context.Context, *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return nil, api.ErrorUserNotFound("user not found")
}

func main() {
	srv := grpc.NewServer(grpc.Address(":9000"))
	helloworld.RegisterGreeterServer(srv, &Service{})

	app := kratos.New(kratos.Server(srv))
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
