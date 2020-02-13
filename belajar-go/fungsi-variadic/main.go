package main

import (
	"fmt"
	"strings"
)

func main() {
	// contoh 1
	// mengisi nilai fungsi variadic dengan cara biasa
	var average = calculate(2, 3, 5, 7, 3, 7, 1, 10)
	var message = fmt.Sprintf("Nilai rata-rata : %.2f", average) // pengunaan Sprintf tidak mengembalikan nilai tapi sebuah string
	fmt.Println(message)

	// mengisi nilai fungsi variadic dengan slice
	var numbers = []int{2, 3, 10, 5, 4, 12}
	var avg = calculate(numbers...)
	var msg = fmt.Sprintf("Nilai rata-rata : %.2f", avg)
	fmt.Println(msg)

	// contoh 2
	// pengisian nilai pada fungsi variadic dengan cara biasa
	myHobbies("andi", "berenang", "footsal")

	// pengisian nilai pada fungsi variadic dengan data slice
	var hobbies = []string{"bola kaki", "badminton"}
	myHobbies("rudi", hobbies...)

}

// contoh 1
func calculate(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}

	var average = float64(total) / float64(len(numbers))
	return average
}

// contoh 2
func myHobbies(name string, hobbies ...string) {
	var joinHobbies = strings.Join(hobbies, ", ")

	fmt.Printf("Hai, my name is : %s\n", name)
	fmt.Printf("My hobbies is : %s\n", joinHobbies)
}
