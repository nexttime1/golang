package main

import (
	"context"
	"fmt"
	"go_GRPC/service_proto/proto"
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
	v_client := proto.NewVideoServiceClient(conn)
	response, err := v_client.Buy(context.Background(), &proto.Request{
		Name: "宋可为❤薛提猛",
	})
	fmt.Println(response, err)
	o_client := proto.NewOrderServiceClient(conn)
	response, err = o_client.Buy(context.Background(), &proto.Request{
		Name: "朝花夕拾Order",
	})
	fmt.Println(response, err)
}
