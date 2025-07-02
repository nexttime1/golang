package main

import (
	"fmt"
	"go_GRPC/stream_proto/proto"
	"google.golang.org/grpc"
	"io"
	"net"
	"os"
)

type StreamService struct {
}

func (StreamService) StreamFunc(request *proto.Request, stream proto.StreamService_StreamFuncServer) error {
	fmt.Println(request)
	// 发送回去  10条消息
	for i := 0; i < 10; i++ {
		stream.Send(&proto.Response{
			Name: fmt.Sprintf("这是我第%d次给你发消息", i),
		})
	}

	return nil
}

func (StreamService) DownLoadFunc(request *proto.Request, stream proto.StreamService_DownLoadFuncServer) error {
	fmt.Println(request)
	// 发送回去一条视频
	file, err := os.Open("static/beijing.mp4")
	if err != nil {
		return err
	}
	defer file.Close()
	for {
		buf := make([]byte, 2048)
		num, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		stream.Send(&proto.ByteResponse{
			Content: buf[:num],
		})
	}

	return nil
}

func (StreamService) UpLoadFile(FileRequest proto.StreamService_UpLoadFileServer) error {
	return nil
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	proto.RegisterStreamServiceServer(server, &StreamService{})
	server.Serve(listen)

}
