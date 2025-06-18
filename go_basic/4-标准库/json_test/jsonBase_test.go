package json_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

type Person struct {
	Name    string
	Age     int
	address string
	married bool
	Parent  []string
}

func TestJson1(t *testing.T) {
	p1 := Person{
		Name:    "xtm",
		Age:     18,
		address: "山东省泰安市岱岳区",
		married: false,
	}
	fmt.Println(p1.address)
	marshal, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(marshal))
	var p2 Person
	err = json.Unmarshal(marshal, &p2)
	fmt.Printf("%#v", p2)
}

// 打开json文件
func TestFIle(t *testing.T) {
	open, err := os.Open("text.json")
	if err != nil {
		fmt.Println(err)
	}
	defer open.Close()
	decoder := json.NewDecoder(open)
	var p4 Person
	err = decoder.Decode(&p4)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", p4)
}

//写到文件区

func TestJsondWriter(t *testing.T) {
	p := Person{
		Name:    "zhangsan",
		Age:     20,
		address: "zhangsan@mail.com",
		Parent:  []string{"Daddy", "Mom"},
	}
	f, _ := os.OpenFile("test.json", os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()

	d := json.NewEncoder(f)
	d.Encode(p)
}
