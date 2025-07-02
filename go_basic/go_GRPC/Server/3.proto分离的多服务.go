package main

import (
	"context"
	"fmt"
	"go_GRPC/service_proto/proto"
	"google.golang.org/grpc"
	"net"
)

type OrderServicer struct {
}

func (OrderServicer) Buy(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	fmt.Println("哈哈哈  调用我了吧 我是OrderServicer")
	return &proto.Response{
		Name: request.Name,
	}, nil
}

type VideoServicer struct {
}

func (VideoServicer) Buy(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	fmt.Println("哈哈哈  调用我了吧 我我我我是VideoServicer")
	return &proto.Response{
		Name: request.Name,
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer() //实例化
	proto.RegisterOrderServiceServer(server, &OrderServicer{})
	proto.RegisterVideoServiceServer(server, &VideoServicer{})
	server.Serve(listen)
	fmt.Println("run 8080")
}
