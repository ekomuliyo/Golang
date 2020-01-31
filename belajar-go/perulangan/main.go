package main

import "fmt"

func main() {

	// menggunakan for
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("")

	// for dengan argumen
	var i = 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// for tanpa argumen
	var j = 0
	for {
		fmt.Println(j)

		j++
		if j == 5 {
			break
		}
	}

	// for dengan break dan continue
	for i := 0; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println(i)
	}

	// for bersarang
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}

	// for memanfaatkan dengan label
outerLoop:
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Printf("matriks [%d][%d] \n", i, j)
		}
	}

}
