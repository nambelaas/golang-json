package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestDecoder(t *testing.T) {
	reader, _ := os.Open("Customer.json")
	decoder := json.NewDecoder(reader)

	customer := &Customer{}

	decoder.Decode(customer)

	fmt.Println(customer)
}
