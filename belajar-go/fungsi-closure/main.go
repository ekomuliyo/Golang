package main

import "fmt"

// Closure adalah sebuah fungsi yang disimpan dalam variabel

func main() {
	// penggunaan closure sederhana
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, element := range n {
			switch {
			case i == 0:
				max, min = element, element
			case element > max:
				max = element
			case element < min:
				min = element
			}
		}
		return min, max
	}

	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 9, 10}
	var min, max = getMinMax(numbers)
	fmt.Printf("data : %v\nmin : %v\nmax : %v\n", numbers, min, max) // template %v digunakan untuk semua jenis tipe data
	fmt.Println()

	// penggunaan closure IIFE (Immediately-Invoked Function Expression) atau memanggil segera ekpresi fungsi
	var numbers2 = []int{2, 4, 6, 3, 7, 8, 3, 10, 14}
	var newNumbers = func(min int) []int {
		var result []int
		for _, element := range numbers2 {
			if element < min {
				continue // jika elemen < dari min maka looping dilanjutkan / melewatkan looping jika elemen < min, kode dibawah nya tidak di eksekusi
			}
			result = append(result, element)
		}
		return result
	}(3) // langsung mengisi nilai argument sebagai parameter dari fungsi ini

	fmt.Println("original number :", numbers2)
	fmt.Println("filtered number :", newNumbers)
	fmt.Println()

	// penggunaan closure sebagai nilai kembalian
	var max2 = 3
	var numbers3 = []int{2, 4, 6, 7, 1, 2, 2, 1, 6}
	var howMany, getNumbers = findMax(numbers3, max2)
	var theNumbers = getNumbers()

	fmt.Println("numbers \t:", numbers3)
	fmt.Printf("value max\t: %d\n\n", max2)

	fmt.Println("length found \t:", howMany)
	fmt.Println("value \t\t:", theNumbers)

}

func findMax(numbers []int, max int) (int, func() []int) {
	var result []int
	for _, element := range numbers {
		if element <= max {
			result = append(result, element)
		}
	}
	return len(result), func() []int {
		return result
	}
}
