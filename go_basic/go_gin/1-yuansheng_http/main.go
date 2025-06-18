package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type DataJson struct {
	Code int                      `json:"code"`
	Name string                   `json:"name"`
	Age  int                      `json:"age"`
	Job  []map[string]interface{} `json:"job"`
}

func Index(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Method)
	fmt.Println(request.URL.String())
	if request.Method != "GET" {
		all, _ := io.ReadAll(request.Body)
		fmt.Println(string(all))
	}
	marshal, _ := json.Marshal(DataJson{
		Code: 200,
		Name: "xtm",
		Age:  18,
		Job:  []map[string]interface{}{{"李白": "2年", "赵云": "3年"}},
	})
	fmt.Println(request.Header)
	writer.Write(marshal)

}

func main() {

	http.HandleFunc("/index", Index)

	http.ListenAndServe("127.0.0.1:8080", nil)

}
