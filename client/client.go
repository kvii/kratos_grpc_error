package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-kratos/examples/errors/api"
	"github.com/go-kratos/examples/helloworld/helloworld"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

func main() {
	ctx := context.Background()
	conn, err := grpc.DialInsecure(ctx, grpc.WithEndpoint("127.0.0.1:9000"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	cli := helloworld.NewGreeterClient(conn)
	_, err = cli.SayHello(ctx, &helloworld.HelloRequest{Name: "kratos"})
	fmt.Printf("Is user not found error? %t\nerror:%v\n", api.IsUserNotFound(err), err)
}
