package main

import (
	"context"
	"fmt"
	"go_GRPC/stream_proto/DoubleProto"
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
	Doubleclient := DoubleProto.NewBothStreamClient(conn)
	chat, err := Doubleclient.Chat(context.Background())
	if err != nil {
		panic(err)
	}
	for i := 0; i < 5; i++ {
		chat.Send(&DoubleProto.Request{
			Content: []byte(fmt.Sprintf("爱的第%d次", i+1)),
		})
		recv, err := chat.Recv()
		if err != nil {
			panic(err)
		}
		fmt.Println(recv)
	}

}
