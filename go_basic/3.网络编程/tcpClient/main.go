package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(os.Stdin) // 创建带缓冲的读取器  读我写的东西
	for {
		readString, err := reader.ReadString('\n') // 读取输入直到遇到换行符
		if err != nil {
			panic(err)
		}
		inputInfo := strings.Trim(readString, "\r\n") //Trim函数的作用是去除字符串两端的指定字符
		if strings.ToUpper(inputInfo) == "Q" {
			break
		}
		_, err = conn.Write([]byte(inputInfo))
		if err != nil {
			panic(err)
		}
		//写过去  我还要读他一会写过来的
		var bytes []byte
		n, err := conn.Read(bytes)
		if err != nil {
			panic(err)
		}
		fmt.Println("接收到服务器给到东西：", string(bytes[:n]))
	}

}
