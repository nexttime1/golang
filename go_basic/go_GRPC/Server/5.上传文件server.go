package main

import (
	"bufio"
	"fmt"
	"go_GRPC/stream_proto/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"os"
)

type Service struct {
}

func (Service) StreamFunc(request *proto.Request, stream proto.StreamService_StreamFuncServer) error {
	fmt.Println(request)
	for i := 0; i < 10; i++ {
		stream.Send(&proto.Response{
			Name:    fmt.Sprintf("第%d次的名字", i),
			Address: fmt.Sprintf("第%d次的地址", i),
		})
	}
	return nil
}

func (Service) DownLoadFunc(request *proto.Request, stream proto.StreamService_DownLoadFuncServer) error {
	fmt.Println(request)
	file, err := os.OpenFile("static/beijing.mp4", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	for {
		buf := make([]byte, 4096)
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		stream.Send(&proto.ByteResponse{
			Content: buf[:n],
		})
	}
	return nil
}

func (Service) UpLoadFile(stream proto.StreamService_UpLoadFileServer) error {
	file, err := os.OpenFile("static/x.png", os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	var index int
	for {
		index++
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		writer.Write(response.Cotent)
		fmt.Printf("第%d次\n", index)
	}
	writer.Flush()
	stream.SendAndClose(&proto.Response{Name: "完毕了"})
	return nil
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	proto.RegisterStreamServiceServer(s, &Service{})
	s.Serve(listen)

}
