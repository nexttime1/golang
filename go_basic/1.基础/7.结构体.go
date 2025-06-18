package main

import (
	"encoding/json"
	"fmt"
)

// 自定义类型
type MyInt_ int

// 起别名
type MyFloat = float64

// 结构体
type Person struct {
	name    string
	age     int
	sex     string
	address string
}

// 结构体的方法
func (p Person) PrintInfo() {
	fmt.Printf("name: %v age: %v\n", p.name, p.age)

}
func (p *Person) SetInfo(name string, age int) {
	p.name = name
	p.age = age

}

// 嵌套结构体
type User struct {
	username string
	age      int
	Address
}
type Address struct {
	street string
	city   string
}

// 继承
type Animal struct {
	Name string
}

type Dog struct {
	Age int
	Animal
}

func (d Dog) wang() {
	fmt.Println("狗叫 汪汪汪\n")
}
func (a Animal) eat() {
	fmt.Printf("%v 在吃饭\n", a.Name)
}

// 嵌套
type Class struct {
	Title    string    `json:"title"` //用于转json 后的重命名
	Students []Student `json:"students"`
}
type Student struct {
	Id  string
	sex string
}

func main() {
	var a MyInt_
	var b MyFloat
	fmt.Printf("a值为%v a的类型是%T\n", a, a) //a值为0 a的类型是main.MyInt_
	fmt.Printf("b值为%v b的类型是%T\n", b, b) // b值为0 b的类型是float64
	var p Person
	p.name = "xtm"
	p.age = 18
	p.sex = "男"
	p.address = "山东"
	fmt.Printf("%#v\n", p) //main.Person{name:"xtm", age:18, sex:"男", address:"山东"}
	//第二种方法
	var p1 = new(Person)
	p1.name = "skw" //本质上是 (*p1).name = 'skw'
	p1.age = 18
	p1.sex = "女"
	p1.address = "<UNK>"
	fmt.Printf("%#v\n", p1) //&main.Person{name:"skw", age:18, sex:"女", address:"<UNK>"}
	//第三种方法
	var p3 = &Person{"libai", 18, "<UNK>", "<UNK>"}
	fmt.Printf("%#v\n", p3) //&main.Person{name:"libai", age:18, sex:"<UNK>", address:"<UNK>"}
	//第四种方法
	var p4 = Person{"libai", 18, "<UNK>", "<UNK>"} //key value 对  key可以省略
	fmt.Printf("%#v\n", p4)                        //main.Person{name:"libai", age:18, sex:"<UNK>", address:"<UNK>"}
	p5 := p3
	p5.name = "111"         //都改  如果不是指针  那就不一样
	fmt.Printf("%#v\n", p5) //&main.Person{name:"111", age:18, sex:"<UNK>", address:"<UNK>"}
	fmt.Printf("%#v\n", p3) //&main.Person{name:"111", age:18, sex:"<UNK>", address:"<UNK>"}

	//结构体中的方法   结构体大写代表公开  小写代表私人
	p5.SetInfo("zhaoyun", 48)
	fmt.Printf("name: %v age: %v \n", p5.name, p5.age)
	p4.PrintInfo()
	u1 := User{
		username: "zhao",
		age:      18,
		Address: Address{
			street: "戴月",
			city:   "泰安",
		},
	}
	u1.Address.street = "代悦"
	fmt.Printf("%#v\n", u1)
	d := Dog{
		Age: 18,
		Animal: Animal{
			Name: "lli",
		},
	}
	d.eat() //调用父结构体的方法
	//结构体转成 json 文件
	marshal, _ := json.Marshal(d) //marshal格式是 byte_ku[]  必须大写才行
	jsonstr := string(marshal)
	fmt.Println(jsonstr)
	//json 转 结构体
	var dd Dog
	err := json.Unmarshal([]byte(jsonstr), &dd)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", dd)
	class1 := Class{
		Title:    "001",
		Students: make([]Student, 0),
	}
	for i := 0; i < 5; i++ {
		student1 := Student{
			Id:  fmt.Sprintf("stu_%d", i),
			sex: "男",
		}
		class1.Students = append(class1.Students, student1)
	}
	fmt.Printf("%#v\n", class1)
	//转成字符串
	bytes, _ := json.Marshal(class1)
	str1 := string(bytes)
	fmt.Println(str1)
	//再转成 结构体
	var class2 = &Class{}
	err = json.Unmarshal([]byte(str1), class2) //由于是指针  不需要加&
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", class2)
}
