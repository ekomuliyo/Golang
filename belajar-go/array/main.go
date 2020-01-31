package main

import "fmt"

func main() {

	// deklarasi array dan setelahnya baru inisialisasi isi
	var nama [5]string
	nama[0] = "andi"
	nama[1] = "rudi"
	nama[2] = "aan"
	nama[3] = "riko"
	nama[4] = "andre"

	fmt.Println(nama)

	fmt.Println()
	// deklarasi array dan lansung mengisi nilainya
	var buah = [3]string{"jeruk", "apel", "mangga"}
	fmt.Println("isi array", buah)
	fmt.Println("ukuran array", len(buah))

	fmt.Println()
	// inisialisasi array di awal tanpa mendefinisikan jumlah array
	var angka = [...]int{1, 2, 3, 4}
	fmt.Println("isi array", angka)
	fmt.Println("ukuran array", len(angka))

	fmt.Println()
	// array multidimensi
	var angka1 = [3][2]int{[2]int{5, 6}, [2]int{1, 2}, [2]int{10, 11}}
	var angka2 = [3][2]int{{1, 2}, {3, 4}, {3, 4}}

	fmt.Println("angka1", angka1)
	fmt.Println("angka2", angka2)

	fmt.Println()
	// mengeluarkan isi array dengan for
	var buah2 = [5]string{"apel", "jeruk", "mangga", "semangka", "pisang"}
	fmt.Println(buah2)

	for i := 0; i < len(buah2); i++ {
		fmt.Printf("\nelement ke %d : %s", i, buah2[i])
	}

	fmt.Println()
	// mengeluarkan isi array dengan for - range
	for i, buah := range buah2 {
		fmt.Printf("\nelement %d : %s", i, buah)
	}

	fmt.Println()
	// menggunakan isi array dengan variabel underscore atau variabel yang tidak berguna
	for _, buah := range buah2 {
		fmt.Printf("\nelement : %s", buah)
	}

	fmt.Println()
	fmt.Println()
	// alokasi element array dengan make([]type_data, panjang_panjang_array)
	var buah3 = make([]string, 3)
	buah3[0] = "apel"
	buah3[1] = "jeruk"
	buah3[2] = "semangka"
	fmt.Println(buah3)
}
