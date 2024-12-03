package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

type Customer struct {
	FirstName string
	LastName  string
	Age       int
	Hobbies   []string
	Addresses []Address
}

func TestJsonObject(t *testing.T) {
	customer := Customer{
		FirstName: "Salman",
		LastName:  "Seif",
		Age:       23,
	}

	bytes, _ := json.Marshal(customer)

	fmt.Println(string(bytes))
}
