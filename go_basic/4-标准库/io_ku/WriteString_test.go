package io_ku

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestWrite(t *testing.T) {
	file, _ := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, 0666)
	defer file.Close()
	n, _ := io.WriteString(file, "hello world  i m golang")
	fmt.Println(n)

}
