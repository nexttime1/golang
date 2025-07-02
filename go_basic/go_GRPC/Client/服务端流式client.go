package main

import (
	"context"
	"fmt"
	"go_GRPC/stream_proto/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	addr := "127.0.0.1:8080"
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := proto.NewStreamServiceClient(conn)
	Result, err := client.StreamFunc(context.Background(), &proto.Request{
		Name: "心比天高",
	})
	if err != nil {
		panic(err)
	}
	for i := 0; i < 10; i++ {
		result, err := Result.Recv()
		if err != nil {
			panic(err)
		}
		fmt.Println(result)
	}
}
