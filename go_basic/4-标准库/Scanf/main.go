package main

import (
	"fmt"
	"strings"
)

func main() {
	var name string
	var age int
	var married bool
	fmt.Scanln(&name, &age, &married)
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(married)
	reader := strings.NewReader("zhangsan 20 true")
	fmt.Fscanf(reader, "%s %d %t\n", &name, &age, &married)
	fmt.Println(age)
}
