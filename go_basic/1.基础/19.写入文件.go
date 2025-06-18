package main

import (
	"bufio"
	"fmt"
	"os"
)

func one() {
	file, err := os.OpenFile("D:\\1111kaoyan111111111111111111111111111111\\go_project\\1.基础\\20_file1.txt", os.O_RDWR|os.O_APPEND, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	file.WriteString("Hello World!\n")
}
func two() {
	file, err := os.OpenFile("D:\\1111kaoyan111111111111111111111111111111\\go_project\\1.基础\\20_file1.txt", os.O_RDONLY|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	newWriter := bufio.NewWriter(file)
	newWriter.WriteString("Hello World,golang!!") //写入缓存种
	newWriter.Flush()                             //写入文件
}
func main() {

	//第一种  写的方法
	one()
	//第二种  写的方法
	two()
}
