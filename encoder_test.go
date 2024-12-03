package belajar_golang_json

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName: "Salman",
		LastName:  "Seif",
		Age:       23,
	}

	encoder.Encode(customer)
}
