package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

// type Customer struct {
// 	Firstname string
// 	Middlename string
// 	Lastname string
// 	Age int
// }

type Person struct {
	Name string
	Age int
}

// Convert JSON to golang type using json.Unmarshal

func TestDecodeJSON(t *testing.T) {
	data := `{"Firstname":"Parar","Middlename":"Tria","Lastname":"Galatian","Age":0}`
	dataBytes := []byte(data)

	customer := &Customer{}
	
	err := json.Unmarshal(dataBytes, customer)
	if err != nil {
		panic(err)
	}


	fmt.Println("data:", data)
	fmt.Println("dataBytes:", dataBytes)
	fmt.Println("customer:", customer)

	personData := `{"Name":"Parar","Age":20}`

	person := &Person{}
	json.Unmarshal([]byte(personData), person)

	fmt.Println("person:", person)
	fmt.Println("Person name:", person.Name)


}