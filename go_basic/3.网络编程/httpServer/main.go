package main

import (
	"fmt"
	"net/http"
)

func main() {
	//注册路由
	http.HandleFunc("/go", handler)
	// 启动HTTP服务器
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic(err)
	}

}
func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("method:", r.Method)
	fmt.Println("url:", r.URL.Path)
	_, err := w.Write([]byte("hello world!! xtm!!!"))
	if err != nil {
		panic(err)
	}
}
