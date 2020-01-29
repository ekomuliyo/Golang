package main

import "fmt"

func main() {

	// operator aritmatika ( *, / , + , - , % )
	var nilai = (((1+2)%2)*10 - 5) / 2
	fmt.Println(nilai)

	// operator perbandingan ( ==, != , < , <= , > , >=)
	var nilai1 = (((1+2)%2)*10 - 5) / 2
	var hasil = (nilai == 2)

	fmt.Printf("nilai %d adalah %t\t\n", nilai1, hasil)

	// operator logika (&& , || , !)
	var kiri = false
	var kanan = true

	fmt.Println("kiri =", kiri)
	fmt.Println("kanan =", kanan)
	var kiriDanKanan = kiri && kanan
	fmt.Printf("kiri && kanan adalah %t \n", kiriDanKanan)

	var kiriAtauKanan = kiri || kanan
	fmt.Printf("Kiri || Kanan adalah %t \n", kiriAtauKanan)

	var kebalikanKiri = !kiri
	fmt.Printf("Kebalikan kiri adalah %t", kebalikanKiri)
}
