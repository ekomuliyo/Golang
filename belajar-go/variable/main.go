package main

import "fmt"

func main()  {

	// deklarasi menggunakan var dan langsung isi variabel nya
	var firstName string = "Budi"
	
	// deklarasi menggunakan var dan mengisinya di lain waktu
	var lastName string
	lastName = "Santoso"

	// definisi variabel langsung menetukan tipe data sesuai isinya
	// atau istilah lain type inference
	fullName := "Andi Santoso"

	// deklarasi multi variabel dan langsung mengisinya menggunakan var
	var satu, dua, tiga string = "one", "two", "three"
	
	empat, lima, enam := "four", "five", "six"

	// variabel underscore atau variabel yang tidak mengembalikan nilai apapun
	// ibarat sebuah lubang, apabila ada sesuatu yang masuk maka akan hilang
	_ = "ini isi variabel underscore tidak dapat ditampilkan"

	// deklarasi menggunakan keyword new
	jurusan := new(string)

	fmt.Printf("hai %s %s, nama lengkap %s\n", firstName, lastName, fullName)

	fmt.Printf("%s %s %s\n %s %s %s\n", satu, dua, tiga, empat, lima, enam)

	fmt.Println(jurusan) // menampilkan hexadesimal
	fmt.Println(*jurusan) // menampilkan string kosong ""
}