package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

func TestJsonTag(t *testing.T) {
	product := Product{
		Id:       "P1",
		Name:     "Apple Macbook Pro",
		ImageUrl: "http://example.com/image.png",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJsonTagDecode(t *testing.T) {
	jsonString := `{"id":"P1","name":"Apple Macbook Pro","image_url":"http://example.com/image.png"}`
	jsonBytes := []byte(jsonString)
	product := &Product{}

	json.Unmarshal(jsonBytes,product)
	fmt.Println(product)
}