package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

func ls_handle(reader *bufio.Reader) { // 改为接收Reader而不是Conn
	// 注意：在handle_accept中已经读取了命令，现在Reader中还有剩余的数据：2|
	num_str, err := reader.ReadString('|')
	if err != nil {
		fmt.Println("读取数字错误", err)
		return
	}
	// 去掉末尾的'|'
	num_str = num_str[:len(num_str)-1]
	int_num, err := strconv.Atoi(num_str)
	if err != nil {
		fmt.Println("转换数字错误", err)
		return
	}
	fmt.Println("我要像你展示的数字为：", int_num)
}

func handle_accept(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	op, err := reader.ReadString('|')
	if err != nil {
		fmt.Println("<UNK>")
		return
	}
	option := op[:len(op)-1]
	fmt.Println("接到的命令是：", option)
	switch option {
	case "ls":
		ls_handle(reader)
	default:
		fmt.Println("无效")
	}
}

func main() {
	file, _ := os.OpenFile("Fileserver.log", os.O_RDWR|os.O_CREATE, 0666)
	defer file.Close()
	//带缓冲的流       这段代码的作用是配置Go的标准日志库，将日志输出重定向到文件"Fileserver.log"。
	log.SetOutput(file)

	addr := "0.0.0.0:9999"
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Printf("listen err %v\n", err)
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("accept err %v\n", err)
			continue
		}
		handle_accept(conn)
	}
}
