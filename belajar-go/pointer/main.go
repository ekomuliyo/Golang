package main

import "fmt"

func main() {

	// contoh 1
	var numberA int = 10
	var numberB *int = &numberA

	fmt.Println("numberA (value)   :", numberA)
	fmt.Println("numberA (address) :", &numberA)
	fmt.Println("numberB (value)   :", *numberB)
	fmt.Println("numberB (address) :", numberB)

	fmt.Println()

	numberA = 8

	fmt.Println("numberA (value)   :", numberA)
	fmt.Println("numberA (address) :", &numberA)
	fmt.Println("numberB (value)   :", *numberB)
	fmt.Println("numberB (address) :", numberB)

	fmt.Println()

	// contoh 2
	var number = 10
	var numberPointer *int = &number

	// sebelum di ubah
	fmt.Println(number)
	fmt.Println(numberPointer)

	// change(number) // tidak berubah karena variabel yang di passing cuman value nya aja dan di fungsi yang menangkap membuat alamat memori baru dengan value yang di passing tsb

	changePointer(&number) // mengubah value variabel nya akan tetapi alamat memori tetap

	fmt.Println(number)
	fmt.Println(numberPointer)

}

// pass by value
func change(value int) {
	value++
}

// pass by reference / pointer
func changePointer(value *int) {
	*value++
}
