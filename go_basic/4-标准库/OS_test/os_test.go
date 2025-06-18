package OS_test

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestCreate(t *testing.T) {

	_, err := os.Create("./test.txt")
	if err != nil {
		t.Fatal(err)
	}
}
func TestOpenFile(t *testing.T) {
	file, _ := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer file.Close()
}

func TestMkdir(t *testing.T) {
	os.Mkdir("ms", 0777)
}

func TestMkdirAll(t *testing.T) {
	os.MkdirAll("ms/a/v/c", 0777)
}

func TestRemove(t *testing.T) {
	os.Remove("ms/a/v/c")
}

func TestRemoveAll(t *testing.T) {
	os.RemoveAll("ms")
}

//获得工作目录

func TestGetwd(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(dir)
}

// 修改当前工作路径
func TestChdir(t *testing.T) {
	err := os.Chdir("D:\\1111kaoyan111111111111111111111111111111\\go_project")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(os.Getwd())
}

// 重命名
func TestRename(t *testing.T) {
	//文件
	err := os.Rename("test.txt", "test001.txt")
	if err != nil {
		t.Fatal(err)
	}
	//目录
	err = os.Rename("ms", "mss")
	if err != nil {
		t.Fatal(err)
	}
}

// readDir
func TestReadDir(t *testing.T) {
	open, _ := os.Open("mss")
	defer open.Close()
	dir, err := open.ReadDir(-1)
	if err != nil {
		t.Fatal(err)
	}
	for _, entry := range dir {
		fmt.Println(entry.Name(), entry.IsDir()) // IsDir  有没有下一级
	}
}

// 偏移读取文件
func TestReadFile(t *testing.T) {
	file, err := os.OpenFile("test001.txt", os.O_RDWR|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	var body []byte
	file.Seek(3, 0) //从第三个字符开始读取
	for {
		cmp := make([]byte, 4)
		n, err := file.Read(cmp)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			t.Fatal(err)
		}
		body = append(body, cmp[:n]...)
		fmt.Println("这次读取了", n, "个字符")
	}
	fmt.Println(string(body))
}

// 文件写
func TestWriteFile(t *testing.T) {
	file, err := os.OpenFile("test001.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	//三种写的方式  第一种
	n, err := file.Write([]byte("谁的爱太疯"))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("写了", n, "字符")
}

// 文件写
func TestWriteFile2(t *testing.T) {
	file, err := os.OpenFile("test001.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	//三种写的方式  第二种   字符串的形式写入
	n, err := file.WriteString("我的爱不疯")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("写了", n, "字符")
}

// 文件写
func TestWriteFile3(t *testing.T) {
	file, err := os.OpenFile("test001.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	//三种写的方式  第三种  指定位置 向后替换
	n, err := file.WriteAt([]byte("什么是爱情"), 5)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("写了", n, "字符")
}
