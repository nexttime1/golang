package main

import (
	"encoding/json"
	"fmt"
)

type Host struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

func (h Host) MarshalJSON() ([]byte, error) {
	var status string
	switch h.Status {
	case Running:
		status = "Running"
	case Except:
		status = "Except"
	case OffLine:
		status = "OffLine"
	}
	return json.Marshal(&struct {
		ID     uint   `json:"id"`
		Name   string `json:"name"`
		Status string `json:"status"`
	}{
		ID:     h.ID,
		Name:   h.Name,
		Status: status,
	})
}

const (
	Running = 1
	Except  = 2
	OffLine = 3
)

func main() {
	host := Host{1, "薛提猛", Running}
	data, _ := json.Marshal(host)
	fmt.Println(string(data)) // {"id":1,"name":"薛提猛","status":"Running"}
}
