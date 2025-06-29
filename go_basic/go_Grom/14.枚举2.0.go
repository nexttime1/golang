package main

import (
	"encoding/json"
	"fmt"
)

type Status int

func (status Status) MarshalJSON() ([]byte, error) {
	var str string
	switch status {
	case Running:
		str = "Running"
	case Except:
		str = "Except"
	case OffLine:
		str = "Status"
	}
	return json.Marshal(str)
}

type Master struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Status Status `json:"status"`
}

func main() {
	host := Master{1, "枫枫", Running}
	data, _ := json.Marshal(host) // 对Status 结构体
	fmt.Println(string(data))     // {"id":1,"name":"枫枫","status":"Running"}
}
