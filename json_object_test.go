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

type User struct {
	Name string
	Age int
	Hobbies []string
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

	// fmt.Println("bytes:", bytes)
	fmt.Println("string(bytes) of customer:", string(bytes))

	user := User{Name: "Parar", Age: 20, Hobbies: []string{"Coding", "Gaming"}}
	userBytes, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	// fmt.Println("userBytes:", userBytes)
	fmt.Println("string(userBytes):", string(userBytes))
	fmt.Println("Hobbies idx 1:", user.Hobbies[1])
}