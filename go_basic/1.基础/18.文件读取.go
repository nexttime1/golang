package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// bufio 去读文件
func read_two() {
	file, _ := os.Open("D:\\1111kaoyan111111111111111111111111111111\\go_project\\1.基础\\20_file.txt")
	reader := bufio.NewReader(file)
	var str string
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			str += line
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		str += line
	}
	fmt.Println(str)
}

func main() {

	open, err := os.Open("D:\\1111kaoyan111111111111111111111111111111\\go_project\\1.基础\\20_file.txt")
	defer open.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	var total_byte []byte
	var str_byte = make([]byte, 128)
	for {
		n, err := open.Read(str_byte)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		total_byte = append(total_byte, str_byte[:n]...)
		fmt.Println("读取了", n, "个字符")
	}
	fmt.Println(string(total_byte))

	//第二种方法
	read_two()

}
