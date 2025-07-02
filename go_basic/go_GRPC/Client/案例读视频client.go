package main

import (
	"bufio"
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
	Result, err := client.DownLoadFunc(context.Background(), &proto.Request{
		Name: "李白",
	})
	if err != nil {
		panic(err)
	}
	file, err := os.OpenFile("static/copy.mp4", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//实现缓冲读
	writer := bufio.NewWriter(file)
	index := 0
	for {
		index++
		result, err := Result.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		writer.Write(result.Content)
		fmt.Printf("这次%d次接受到的数据，共%d byte\n", index, len(result.Content))

	}

}
