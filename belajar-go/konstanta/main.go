package main

import "fmt"

func main() {

	// konstanta biasnya digunakan untuk nilai yang tidak dapat diubah
	// penggunaan variabel dengan tipe data
	const nama string = "Eko Muliyo"

	fmt.Println(nama)

	// penggunaan variabel dengan type inference atau tipe data datanya didefinisika otomasi sesuai isinya
	const namaPanggilan = "Eko"
	fmt.Println(namaPanggilan)

}
