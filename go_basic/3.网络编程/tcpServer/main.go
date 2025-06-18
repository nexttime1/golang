package main

import (
	"bufio"
	"fmt"
	"net"
)

func procress(conn net.Conn) {
	defer conn.Close()
	for {
		var buf [1024]byte
		reader := bufio.NewReader(conn) //创建缓冲读取器后再读取
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("end  结束")
			break
		}
		v := string(buf[:n])
		fmt.Println("收到客户端的信息是", v)
		_, err = conn.Write([]byte(v))
		if err != nil {
			panic(err)
		}
	}
}

func main() {

	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		go procress(conn)
	}
}
