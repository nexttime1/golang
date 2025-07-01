package main

import (
	"context"
	"fmt"
	"go_GRPC/duo_grpc"
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

	videoClient := duo_grpc.NewVideoServerClient(conn)
	res, err := videoClient.Look(context.Background(), &duo_grpc.VideoRequest{
		Name: "李白",
	})
	fmt.Println(res)
	orderClient := duo_grpc.NewOrderServerClient(conn)
	result, err := orderClient.Buy(context.Background(), &duo_grpc.OrderRequest{
		Name: "猴子",
	})
	fmt.Println(result)

}
