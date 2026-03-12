package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id string `json:"id"`
	Name string `json:"name"`
	ImageUrl string `json:"image_url"`
}

func TestJsonTagEncode(t *testing.T) {

	product := Product{
		Id: "P001",
		Name: "Macbook Pro M3",
		ImageUrl: "https://example.com/macbook_pro_m3.jpg",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJsonTagDecode(t *testing.T) {
	jsonData := `{"id":"P001","name":"Macbook Pro M3","image_url":"https://example.com/macbook_pro_m3.jpg"}`
	jsonBytes := []byte(jsonData)

	product := &Product{}

	err := json.Unmarshal(jsonBytes, product)
	if err != nil {
		panic(err)
	}

	fmt.Println("Product:", product)
}