package main

import (
	"fmt"
	"runtime"
)

func main() {

	// perbedaan konkurensi dengan pararel
	// konkurensi adalah komposisi dari sebuah proses
	// pararel adalah mengjalankan proses secara bersamaan
	// goroutine adalah salah satu konkurensi di bahasa Go

	runtime.GOMAXPROCS(2)

	go print(5, "halo")   // memanggil fungsi print dengan menerapkan goroutine
	print(5, "apa kabar") // memanggil fungsi dengan cara biasa

	// menghandle agar proses yang dijalankan di goroutine dapat di eksekusi, proses jalannya aplikasi berhenti di baris tsb (blocking)
	// karena eksekusi program di atas di 2 proses, dan ketika proses goroutine (go print(5, "halo")) lebih lama dibading proses utama (print(5 "apakabar"))
	// maka proses goroutine tidak terkesekusi, karena program di anggap telah selesai
	var input string
	fmt.Scanln(&input)

	// penggunaan Scanln()
	// mengcapture semua karakter setelah user menekan enter / input
	var a, b, c string
	fmt.Scanln(&a, &b, &c)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}
