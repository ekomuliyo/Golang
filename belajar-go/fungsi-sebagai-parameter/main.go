package main

import (
	"fmt"
	"strings"
)

func main() {
	var data = []string{"Andi", "Budi sudarsono", "Rudi saputra"}

	var containsA = filter(data, func(elemen string) bool {
		return strings.Contains(elemen, "A") // strings.Contains(data, "search") mengembailkan nilai true jika terdapat huruf "search" dalam data
	})

	var dataLength = filter(data, func(elemen string) bool {
		return len(elemen) == 14
	})

	fmt.Println("data original \t\t:", data)
	fmt.Println("filter huruf \"A\" \t:", containsA)
	fmt.Println("filter jumlah huruf \"14\" :", dataLength)
}

// FilterCallback adalah merepresentasikan FilterCallback
type FilterCallback func(string) bool

func filter(data []string, callback FilterCallback) []string {
	var result []string

	for _, elemen := range data {
		if filtered := callback(elemen); filtered { // parameter callback akan dipanggil sesuai looping nya dan method tersebut memberikan nilai sesuai dengan yang return nya
			result = append(result, elemen)
		}
	}
	return result
}
