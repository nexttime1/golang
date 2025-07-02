package main

import (
	"context"
	"fmt"
	"go_GRPC/stream_proto/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"os"
)

func main() {

	addr := "127.0.0.1:8080"
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := proto.NewStreamServiceClient(conn)
	Result, err := client.UpLoadFile(context.Background())
	if err != nil {
		panic(err)
	}
	file, err := os.OpenFile("static/linwan.png", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	index := 0
	for {
		index++
		buf := make([]byte, 2048)
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		Result.Send(&proto.FileRequest{
			Cotent: buf[:n],
		})
		fmt.Printf("第%d次", index)
	}

}
