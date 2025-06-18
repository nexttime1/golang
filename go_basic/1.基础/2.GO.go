package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	//条件
	age := 18
	if age < 18 {
		fmt.Println("未成年")
	} else {
		fmt.Println("成年")
	}
	if Age := 20; Age > 10 {
		fmt.Println("大于10岁") //局部变量 Age

	}
	//for 三种写法
	for i := 0; i < 1; i++ { //i也是局部变量
		fmt.Println(i)
	}
	i := 1
	for ; i < 2; i++ {
		fmt.Println(i)
	}
	for i < 5 {
		fmt.Println(i)
		i++
	}
	//打印
	row := 5
	for i := 1; i <= row; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
	// 9 9 乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d  ", j, i, i*j)
		}
		fmt.Println("")
	}
	//switch  可以多个值  可以不用break
	score := "A"
	switch score {
	case "A", "B", "C":
		fmt.Println("及格")
	case "D":
		fmt.Println("不及格")
	}
	age = 30
	switch {
	case age < 24:
		fmt.Println("好好学习")
	case age > 24 && age < 60:
		fmt.Println("好好工作")
		fallthrough //贯穿 下一个
	case age > 60:
		fmt.Println("注意身体")
	default:
		fmt.Println("输入错误")
	}
	//break跳多个循环
label:
	for i := 0; i < 5; i++ {
		for j := 1; j <= 5; j++ {
			if j == 3 {
				break label //跳两个循环  直接结束
			}
			fmt.Println(i, j)
		}
	}
label1:
	for i := 0; i < 5; i++ {
		for j := 1; j <= 5; j++ {
			if j == 3 {
				continue label1 //continue 到 上面  继续执行
			}
			//fmt.Println(i, j)
		}
	}
	//goto
	for i := 0; i < 5; i++ {
		if i == 3 {
			goto label2
		}
	}
	fmt.Println("aaa")
label2:
	fmt.Println("gogogo")

	//数组
	var str = [3]string{"java", "golang", "python"}
	s := str[0]
	list_tune := []rune(s)
	list_tune[0] = 74
	//fmt.Println(string(list_tune)) //Java
	str[0] = string(list_tune)
	fmt.Println(str)
	var a = [...]int{1, 2, 3, 4, 5, 6} //电脑自己数
	fmt.Println(a)
	var b [5]string
	fmt.Println(b) //默认为空   int 默认为0
	//数组是值类型  切片是引用
	var a1 = [...][2]string{
		{"北京", "上海"},
		{"济南", "青岛"},
	}
	var a2 = a1
	a2[0][0] = "深圳"
	fmt.Println(a2)
	fmt.Println(a1)
	var c1 = []string{"<UNK>", "<UNK>", "<UNK>", "<UNK>"}
	c2 := c1
	c1[0] = "xtm"
	fmt.Println(c1) //都改变
	fmt.Println(c2) //都改变
	//切片 指针  cap 容量 底层数组 一直到最后的那个数
	split_str := []string{"1", "2", "3", "4", "5", "6"}
	string_split1 := split_str[:4]
	fmt.Printf("长度为:%d, 容量为: %d\n", len(string_split1), cap(string_split1)) //长度为:4, 容量为: 6
	//make
	var string_split2 = make([]string, 4)                   //长度是4
	string_split2 = append(string_split2, string_split1...) //固定搭配
	string_split2 = append(string_split2, "1123", "223")
	string_split2[0] = "xtm"
	fmt.Println(string_split2, len(string_split2))
	//copy
	split_str1 := []int{1, 2, 3, 4, 5, 6}
	split_str2 := make([]int, len(split_str1))
	copy(split_str2, split_str1)
	split_str2[0] = 111
	fmt.Println(split_str2)
	fmt.Println(split_str1)
	//删除下标为2的元素
	split_str1 = append(split_str1[:2], split_str1[3:]...)
	fmt.Println(split_str1)

	//map
	userinfo := make(map[string]string)
	userinfo["name"] = "xtm"
	userinfo["age"] = "20"
	userinfo1 := map[string]string{
		"name": "xtm",
		"age":  "20",
	}
	fmt.Println(userinfo1)
	//遍历   和  取出值  删除一对
	for k, v := range userinfo {
		fmt.Printf("键为： %v 值为：%v\n", k, v)
	}
	value, ok := userinfo["name"]
	fmt.Println(value, ok) //有就赋值  ok = true  没有就 空  ok = false
	userinfo["sex"] = "男"
	delete(userinfo, "age")
	fmt.Println(userinfo)

	//一些用户信息
	userinfo_some := make([]map[string]string, 2)
	userinfo_some[0] = make(map[string]string)
	userinfo_some[0]["name"] = "xtm"
	userinfo_some[0]["age"] = "20"
	userinfo_some[1] = make(map[string]string)
	userinfo_some[1]["name"] = "zhaoyun"
	userinfo_some[1]["age"] = "40"
	fmt.Println(userinfo_some)
	for _, v := range userinfo_some {
		for k1, v1 := range v {
			fmt.Println(k1, v1)
		}
	}
	//value 是切片时
	userinfo3 := make(map[string][]string)
	userinfo3["hobby"] = []string{"玩游戏", "敲代码", "旅游"}
	fmt.Println(userinfo3)

	//按key升序输出 字典
	dict1 := map[int]int{
		1:  20,
		8:  9,
		4:  13,
		20: 22,
		7:  99,
	}
	split1 := make([]int, 0)
	for k, _ := range dict1 {
		split1 = append(split1, k)
	}
	fmt.Println(split1)
	sort.Ints(split1)
	for _, v := range split1 {
		fmt.Printf("key为：%v  value为：%v\n", v, dict1[v])
	}
	//统计 how do you do  单词出现的个数
	strr1 := "how do you do"
	strSlice := strings.Split(strr1, " ")
	fmt.Println(strSlice)
	dict2 := make(map[string]int)
	for _, v := range strSlice {
		_, ok := dict2[v]
		if ok {
			dict2[v]++
		} else {
			dict2[v] = 1
		}
	}
	fmt.Println(dict2)
}
