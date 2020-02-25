package main

import (
	"fmt"
	"strings"
)

func main() {
	// penggunaan interface
	var secret interface{}

	secret = "eko"
	fmt.Println(secret)

	fmt.Println()
	secret = []string{"jeruk", "apel", "semangka"}
	fmt.Println(secret)

	fmt.Println()
	secret = 12.4
	fmt.Println(secret)

	fmt.Println()
	var data map[string]interface{}
	data = map[string]interface{}{
		"name":      "eko muliyo",
		"grade":     20,
		"breakfast": []string{"apel", "jeruk", "semangka"},
	}
	fmt.Println(data)

	fmt.Println()
	// casting variabel interface agar bisa operasikan
	secret = 2
	var number = secret.(int) * 20
	fmt.Println(secret, "hasil setelah di casting", number)

	fmt.Println()
	secret = []string{"apel", "jeruk", "semangka"}
	var fruits = strings.Join(secret.([]string), ", ")
	fmt.Println(fruits, "adalah sebuah hasil dari join string array")

	fmt.Println()
	// casting variabel interface kosong ke object pointer
	type person struct {
		name string
		age  int
	}
	secret = &person{name: "budi", age: 19}
	var name = secret.(*person).name
	fmt.Println(name)

	fmt.Println()
	// kombinasi slice, map, interface{}
	var person2 = []map[string]interface{}{
		{"name": "andi", "age": 22},
		{"name": "budi", "age": 21},
		{"name": "aan", "age": 20},
	}

	for _, each := range person2 {
		fmt.Println(each["name"], "age is", each["age"])
	}

	// penggunaan slice dengan interface{}, kita dapat menampung nilai type apapun
	var fruits2 = []interface{}{
		map[string]interface{}{
			"name": "andi", "umur": 19,
		},
		[]string{"pempek", "model", "tekwan"},
		"jeruk",
	}

	fmt.Println()
	for _, each := range fruits2 {
		fmt.Println(each)
	}
}
