package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	udp, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8880,
	})
	if err != nil {
		panic(err)
	}
	defer udp.Close()
	readString, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		panic(err)
	}

	_, err = udp.Write([]byte(readString))
	if err != nil {
		panic(err)
	}
	buf := make([]byte, 1024)
	n, addr, err := udp.ReadFromUDP(buf)
	if err != nil {
		panic(err)
	}
	fmt.Println("接受到服务器的数据", string(buf[:n]), "地址来源：", addr)

}
