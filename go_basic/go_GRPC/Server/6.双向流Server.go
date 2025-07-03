package main

import (
	"fmt"
	"go_GRPC/stream_proto/DoubleProto"
	"google.golang.org/grpc"
	"net"
)

type BothStream struct {
}

func (BothStream) Chat(stream DoubleProto.BothStream_ChatServer) error {
	for i := 0; i < 5; i++ {
		res, _ := stream.Recv()
		fmt.Println(string(res.Content))
		stream.Send(&DoubleProto.Response{
			Content: []byte(fmt.Sprintf("第%d次发给你", i+1)),
		})
	}
	return nil
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	DoubleProto.RegisterBothStreamServer(s, &BothStream{})
	s.Serve(listen)

}
