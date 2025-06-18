package main

import (
	"fmt"
	"io"
	"os"
)

func copy1(scrName string, DesMame string) (err error) {
	file, e1 := os.Open(scrName)
	openFile, e2 := os.OpenFile(DesMame, os.O_RDWR|os.O_CREATE, 0666)
	defer file.Close()
	defer openFile.Close()
	if e1 != nil {
		panic(e1)
		return e1
	}
	if e2 != nil {
		panic(e2)
		return e2
	}
	bytes := make([]byte, 1024)
	for {
		n1, e3 := file.Read(bytes)
		if e3 == io.EOF {
			break
		} else if e3 != nil {
			panic(e3)
			return e3
		}
		_, e4 := openFile.Write(bytes[:n1])
		if e4 != nil {
			return e4
		}
	}
	return nil
}

func main() {
	src := "D:\\1111kaoyan111111111111111111111111111111\\go_project\\1.基础\\20_file1.txt"
	desName := "D:\\1111kaoyan111111111111111111111111111111\\go_project\\1.基础\\20_file1111.txt"
	err := copy1(src, desName)
	if err == nil {
		fmt.Println("复制成功")
	}
	//当前位置创建目录

	os.Mkdir("./abc", 0666)

	//创建多级目录
	os.MkdirAll("./abc/a/v/b/d", 0666)

	//删除一个文件
	os.Remove("./text.txt")
	//删除多个文件
	os.RemoveAll("./abc") //目录下全部删除
	//重命名
	os.Rename("./abc/1.txt", "./abc/text.txt")

}
