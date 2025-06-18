package io_ku

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"testing"
)

func TestPipe(t *testing.T) {
	r, w := io.Pipe()

	go func() {
		fmt.Fprint(w, "some io.Reader stream to be read\n")
		w.Close()
	}()

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}

}

func TestReadAll(T *testing.T) {
	//r := strings.NewReader("Go is a general-purpose language designed with systems programming in mind.")
	r1 := bytes.NewReader([]byte("Go is a general-purpose language designed with systems programming in mind."))

	b, err := io.ReadAll(r1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", b)
}
