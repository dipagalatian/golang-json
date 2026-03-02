package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	Firstname string
	Middlename string
	Lastname string
	Age int
}

// Convert golang type to JSON Object using json.Marshal func
func TestJSONObject(t *testing.T) {
	customer := Customer{
		Firstname: "Parar",
		Middlename: "Tria",
		Lastname: "Galatian",
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println("bytes:", bytes)
	fmt.Println("string(bytes):", string(bytes))
}