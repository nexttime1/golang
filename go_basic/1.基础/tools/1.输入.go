package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	Text := scanner.Text()
	fmt.Printf("打印的是 %s\n", Text)
	split := strings.Split(Text, " ")
	for _, v := range split {
		fmt.Printf("<UNK> %s,类型是 %T\n", v, v)
	}
}
