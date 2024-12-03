package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName: "Salman",
		LastName:  "Seif",
		Age:       23,
		Hobbies:   []string{"Gaming", "Reading", "Coding"},
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Salman","LastName":"Seif","Age":23,"Hobbies":["Gaming","Reading","Coding"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)

	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Salman",
		Addresses: []Address{
			{
				Street: "Jalan ini",
				Country: "Indonesia",
				PostalCode: "123",
			},
			{
				Street: "Jalan itu",
				Country: "Indonesia",
				PostalCode: "456",
			},
		},
	}

	bytes,_ := json.Marshal(customer)

	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Salman","LastName":"","Age":0,"Hobbies":null,"Addresses":[{"Street":"Jalan ini","Country":"Indonesia","PostalCode":"123"},{"Street":"Jalan itu","Country":"Indonesia","PostalCode":"456"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)

	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Addresses)
}

func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Jalan ini","Country":"Indonesia","PostalCode":"123"},{"Street":"Jalan itu","Country":"Indonesia","PostalCode":"456"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}

	err := json.Unmarshal(jsonBytes, addresses)

	if err != nil {
		panic(err)
	}

	fmt.Println(addresses)
}

func TestOnlyJSONArrayComplex(t *testing.T) {

	addresses := []Address{
			{
				Street: "Jalan ini",
				Country: "Indonesia",
				PostalCode: "123",
			},
			{
				Street: "Jalan itu",
				Country: "Indonesia",
				PostalCode: "456",
			},
		}

	bytes, err := json.Marshal(addresses)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}