package main

import (
	"context"
	"fmt"
	"go_GRPC/duo_grpc"
	"google.golang.org/grpc"
	"net"
)

type VideoService struct {
}

func (VideoService) Look(ctx context.Context, Request *duo_grpc.VideoRequest) (*duo_grpc.VideoResponse, error) {
	fmt.Println("Look")
	return &duo_grpc.VideoResponse{
		Name: "薛提猛",
	}, nil
}

type OrderService struct {
}

func (OrderService) Buy(ctx context.Context, Request *duo_grpc.OrderRequest) (*duo_grpc.OrderResponse, error) {
	fmt.Println("order")
	return &duo_grpc.OrderResponse{
		Name: "宋可为",
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	duo_grpc.RegisterVideoServerServer(server, &VideoService{})
	duo_grpc.RegisterOrderServerServer(server, &OrderService{})
	fmt.Println("Server is running on port :8080")
	server.Serve(listen)

}
