package main

import "fmt"

func main()  {

	// menetukan tipe data int8 dengan rentang angka 0 - 255
	var angkaPositif uint8 = 85

	// type data otomatis ditentukan sesuai nilainya
	var angkaNegatif = -78927958

	fmt.Printf("bilangan positif : %d\n", angkaPositif)
	fmt.Printf("bilangan negatif : %d\n", angkaNegatif)

	// type data float ada dua yaitu float32 dan float64
	var angkaDesimal = 26.42
	fmt.Printf("bilangan desimal : %f\n", angkaDesimal)
	fmt.Printf("bilangan desimal, 3 angka dibelakang koma : %.3f\n", angkaDesimal)

	// type data boolean
	var kondisi bool = true

	fmt.Printf("Kondisi : %t \n", kondisi)

	// type data string
	var pesan string = "ini contoh text menggunakan type data string"
	// menggunakan backticks agar bisa digunakan jika text nya lebih dari 1 baris
	var pesanLagi string = `ini contoh text string menggunakan backticks
	agar bisa dua baris seperti ini,
	keren kan`

	fmt.Printf("text : %s\n", pesan)
	fmt.Printf("text backticks : %s\n", pesanLagi)

	// konstanta, sebuah variabel yang mempunyai nilai yang tidak bisa di ubah
	const nama = "andi"
	fmt.Printf("nama saya %s \n", nama)
}