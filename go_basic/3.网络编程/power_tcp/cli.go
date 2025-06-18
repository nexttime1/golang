package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

/*
fmt.Fprintf内部会格式化字符串（这里没有额外的参数，所以就是原字符串）。
然后调用conn.Write([]byte_ku)方法将字节数据写入连接。
操作系统会将这个数据通过TCP（或其他协议）发送到对端。
*/

func ls(conn net.Conn, str string) {
	fmt.Fprintf(conn, "ls|"+str+"|")
}

func main() {
	addr := "127.0.0.1:9999"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	Test := reader.Text()
	message := strings.Split(Test, " ")
	switch message[0] {
	case "ls":
		ls(conn, message[1])
	}
}
