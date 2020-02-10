package main

import "fmt"

func main() {
	var bulan = map[string]int{}

	bulan["januari"] = 1
	bulan["februari"] = 2

	fmt.Println("januari =", bulan["januari"])
	fmt.Println("februari =", bulan["februari"])
	fmt.Println("maret =", bulan["maret"]) // apabila mengakses key dari map yang tidak ada nilainya maka akan mengembalikan nilai default, dalam hal ini default int 0

	// inisialisasi nilai map
	// inisialisasi menggunakan new, make
	// var bulan2 = new(map[string]int) ini akan menghasilkan data pointer
	// var bulan3 = make(map[string]int)

	var bulan1 = map[string]int{
		"januari":  1,
		"februari": 2,
		"maret":    3,
	}

	for key, value := range bulan1 {
		fmt.Println(key, "  \t:", value)
	}

	// menghapus item map
	var bulan2 = map[string]int{"april": 4, "mei": 5, "juni": 6}
	fmt.Println(bulan2)
	delete(bulan2, "april")
	fmt.Println(bulan2)

	// deteksi keberadaan item dalam map
	var bulan3 = map[string]int{"juli": 7, "agustus": 8, "september": 9}
	var value, isExist = bulan3["juli"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("no found")
	}

	// kombinasi slice dan map
	var bulan4 = []map[string]string{
		{"name": "januari", "number": "satu", "english": "one"},
		{"name": "februari", "number": "dua", "id": "2"},
		{"name": "maret", "number": "tiga", "english": "three"},
	}

	for _, value := range bulan4 {
		fmt.Println(value["name"], value["number"], value["id"])
	}
}
