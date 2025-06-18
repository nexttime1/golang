package main

import (
	"fmt"
	"strconv"
	"strings"
	"unsafe"
)

// ()形参   第二个括号代表返回
func Get_Info() (string, int) {
	return "xtm", 20
}
func test() bool {
	fmt.Println("test")
	return true
}

func main() {
	var (
		username = "xtm"
		sex      = "男"
		age      = 20
	)
	fmt.Printf("Hello %v ,我记得你的性别是%v ，年龄是%d \n", username, sex, age)

	a, b, c := 1, 2, "skw"
	fmt.Println(a, b, c)
	//匿名变量 _
	var name, _ = Get_Info()
	fmt.Println(name)

	//	常量
	const PI = 3.14159
	fmt.Println(PI)
	const (
		n1 = iota
		n2
		n3
	)
	fmt.Println(n1, n2, n3)

	//unsafe
	var Age int8 = 20
	fmt.Println(Age)
	fmt.Println(unsafe.Sizeof(Age)) //字节数
	//强制类型转换
	var number1 int16 = 111
	var number2 int32 = 222
	fmt.Println(number1 + int16(number2))
	fmt.Printf("number1 原样输出是 %v \n", number1)
	fmt.Printf("number2 十进制输出是 %d \n", number1)
	fmt.Printf("number2 二进制输出是 %b \n", number1)
	fmt.Printf("number2 16进制输出是 %x \n", number1)
	//浮点型
	var f1 float64 = 3.14159
	var f2 float32 = 4.2222
	fmt.Println(f1 + float64(f2))
	fmt.Printf("f1 == %.2f \n", f1)
	//精度丢失
	var m1 float32 = 8.2
	var m2 float32 = 3.8
	fmt.Println(m1 - m2) //4.3999996

	//bool
	var b1 bool = true
	fmt.Println(b1)
	if b1 {
		fmt.Println("b1 is true")
	} else {
		fmt.Println("b1 is false")
	}

	//字符串
	var str1 = "this is \n str"
	fmt.Println(str1)
	var str2 = "GO\\bin" // GO\bin   转译
	fmt.Println(str2)
	str2 = "go\"bin" //go"bin
	fmt.Println(str2)
	//拼接字符串
	str3 := "xtm"
	str4 := "skw"
	fmt.Println(str3 + str4)
	str5 := fmt.Sprintf("%v 喜欢 %v", str3, str4)
	fmt.Println(str5)
	//切片
	str6 := "123-456-789"
	var arr = strings.Split(str6, "-")
	fmt.Println(arr) //[123 456 789]
	//合成str
	str6 = strings.Join(arr, "")
	fmt.Println(str6) //123456789
	//是否包含
	str1 = "this is xtm"
	str2 = "this"
	flag := strings.Contains(str1, str2)
	fmt.Println(flag)                    //true
	flag = strings.HasPrefix(str1, str2) //前缀是否是
	fmt.Println(flag)                    //true
	flag = strings.HasSuffix(str1, str2) //后缀是否是
	fmt.Println(flag)                    //false

	//查找下标
	str1 = "this is xtm"
	str2 = "is"
	num := strings.Index(str1, str2)    //从前向后找
	fmt.Println(num)                    //2
	num = strings.LastIndex(str1, str2) //从后向前
	fmt.Println(num)                    //5
	//单个字符 是 int
	var c1 = 'a'
	fmt.Printf("原样输出 %c  输出：%v 类型是%T \n", c1, c1, c1) // 97
	c2 := "this is xtm"
	for i := 0; i < len(c2); i++ {
		fmt.Printf("%c(%v)", c2[i], c2[i])
	}
	for i1, v := range c2 {
		fmt.Printf("\n%c(%v)", v, v)
		fmt.Printf("下标是%v", i1)
	}
	//字符串是不可变类型  要改变  必须转成  字节和rune类型 字节是ASCII  rune是汉字
	str7 := "abc"
	bytes := []byte(str7)
	bytes[0] = 'x'
	fmt.Println(string(bytes)) //xbc
	str8 := "我们的歌"
	runes := []rune(str8)
	runes[0] = '你'
	fmt.Println(string(runes)) //你们的歌
	//进制转换
	strr1 := 12
	strr2 := 12.5
	strr3 := true
	strr4 := 'a'
	// %d int  %c string/byte_ku  %t bool  %f float
	strr5 := fmt.Sprintf("%d", strr1)
	fmt.Printf("%v 类型是%T\n", strr5, strr5)
	strr5 = fmt.Sprintf("%.2f", strr2)
	fmt.Printf("%v 类型是%T\n", strr5, strr5)
	strr5 = fmt.Sprintf("%t", strr3)
	fmt.Printf("%v 类型是%T\n", strr5, strr5)
	strr5 = fmt.Sprintf("%c", strr4)
	fmt.Printf("%v 类型是%T\n", strr5, strr5)
	//用包 把其他类型转换成 str
	numb1 := 12
	numb2 := 3.1415
	strr6 := strconv.FormatInt(int64(numb1), 10)
	fmt.Printf("%v 类型是%T\n", strr6, strr6)
	strr6 = strconv.FormatFloat(float64(numb2), 'f', 3, 64) //保留3位小数
	fmt.Printf("%v 类型是%T\n", strr6, strr6)
	//字符转换成 int  和  float
	str1 = "123"
	str2 = "44.55"
	num1, _ := strconv.ParseInt(str1, 10, 64)
	//num1, err := strconv.Atoi(str1)  转换为整数
	float1, _ := strconv.ParseFloat(str2, 64)
	fmt.Printf("%v 类型是%T\n", num1, num1)
	fmt.Printf("%v 类型是%T\n", float1, float1)
	numb1 = 2
	numb1++
	flag = numb1 > 10 && test() //不执行 test
	fmt.Println(flag)
	//位运算
	numb1 = 5                  // 0101
	numb5 := 2                 // 0010
	fmt.Println(numb1 & numb5) //0
	fmt.Println(numb1 | numb5) //7
	fmt.Println(numb1 ^ numb5) //异或  7
	fmt.Println(numb5 << 3)    //16
}
