package main

import (
	"context"
	"fmt"
	"go_GRPC/hello_grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	addr := "127.0.0.1:8080"
	// 必须显式指定安全选项，强制开发者考虑安全性
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := hello_grpc.NewHelloServiceClient(conn) //初始化客户端
	response, err := client.SayHello(context.Background(), &hello_grpc.HelloRequest{
		Name:    "宋可为",
		Message: "爱薛提猛到无法自拔",
	})
	fmt.Println(response, err)

}
