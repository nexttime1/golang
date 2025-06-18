package scorce

import "fmt"

type Roman struct {
	Name string
	Age  int
}

func (r Roman) GetName() {
	fmt.Println(r.Name)
}
