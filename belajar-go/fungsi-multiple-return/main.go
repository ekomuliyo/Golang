package main

import (
	"fmt"
	"math"
)

func main() {
	var diameter float64 = 15

	// multiple return
	var area, circumference = calculate(diameter)
	fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
	fmt.Printf("keliling lingkaran\t: %.2f \n", circumference)

	fmt.Println()

	// multiple return predefined
	var a, c = calculatePredefined(diameter)
	fmt.Printf("luas lingkaran\t\t: %.2f \n", a)
	fmt.Printf("keliling lingkaran\t: %.2f \n", c)
}

// multiple return
func calculate(d float64) (float64, float64) {
	// hitung luas
	// math.Pow adalah pangkat contoh math.Pow(a,b) = a^b
	// math.Pi adalah bilang Pi = 22/7
	var area = math.Pi * math.Pow(d/2, 2)

	// hitung keliling
	var circumference = math.Pi * d

	// kembalikan 2 nilai
	return area, circumference
}

// multriple return dengan predefined
// mengembalikan nilai sesuai dengan nama variabel yang didefinisikan diawal

func calculatePredefined(d float64) (a float64, c float64) {
	a = math.Pi * math.Pow(d/2, 2)
	c = math.Pi * d

	return
}
