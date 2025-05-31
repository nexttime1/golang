package golang_test

import (
	"fmt"
	//"github.com/nexttime1/golang"
	"golang"
	"testing"
)

func TestHello(t *testing.T) {
	data := "jack"
	expected := fmt.Sprintf("hello %s!", data)
	result := golang.Hello(data)

	if result != expected {
		t.Fatalf("expected result %s, but got %s", expected, result)
	}
}
