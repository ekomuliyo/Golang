package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	var names = []string{"Eko", "Muliyo"}
	printMessage("hai", names)

	rand.Seed(time.Now().Unix())
	var randomValue int
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number :", randomValue)

	divideNumber(10, 0)
}

// fungsi sederhana
func printMessage(message string, ar []string) {
	var nameString = strings.Join(ar, " ")
	fmt.Println(message, nameString)
}

// fungsi dengan nilai balik / return
// disini type parameter semua bernilai int, maka cukup dengan mendeklarasikan type di parameter akhir
func randomWithRange(min, max int) int {
	var value = rand.Int()%(max-min+1) + min
	return value
}

// penggunaan return untuk menghentikan suatu proses dalam fungsi
func divideNumber(m, n int) {
	if m == 0 {
		fmt.Printf("invalid divider. %d cannot divided by %d", m, n)
		return
	} else if n == 0 {
		fmt.Printf("invalid divider. %d cannot divided by %d", n, m)
		return
	}

	var res = m / n
	fmt.Printf("%d / %d = %d\n", m, n, res)
}
