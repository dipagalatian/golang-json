package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

// json.Marshal is used for encode golang type to json type
// json.Unmarshal is used for decode json type to golang type
func logJson(data any) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
	fmt.Printf("Type data: %T\n", data)
}

func LoginUser(data any) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	dataStr := string(bytes)
	fmt.Println("Login data: ", dataStr)
}

func TestEncode(t *testing.T) {
	logJson("abc")
	logJson(1)
	logJson(true)
	logJson([]string{"satu", "dua", "tiga"})

	LoginUser(map[string]string{
		"id": "1",
		"name": "dipa",
		"email": "dipa@gmail.com",
	})
}