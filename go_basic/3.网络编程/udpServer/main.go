package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8880,
	})
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	buf := make([]byte, 1024)
	for {
		n, addr, err := listener.ReadFromUDP(buf)
		if err != nil {
			panic(err)
		}
		fmt.Print("收到客户端的消息：", string(buf[:n]))
		fmt.Println("   地址为：", addr.String())
		//我在写回去
		_, err = listener.WriteToUDP(buf[:n], addr)
		if err != nil {
			panic(err)
		}
	}
}
