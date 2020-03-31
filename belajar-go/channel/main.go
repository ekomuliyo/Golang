package main

import (
	"fmt"
	"runtime"
)

// channel digunakan untuk menghubungkan antara satu goroutine dengan goroutine lainnya
// mekanisme channel sederhana, yaitu ada channel pengirim dan channel penerima data
// dan mekanisme ini bersifat blocking / synchronous

func main() {

	// Penerapan Channel

	// mengatur goroutine dengan 2 prosessor yang akan dijalankan
	runtime.GOMAXPROCS(2)

	// deklarasi variabel channel dengan tipe data argumen string
	var messageChannel1 = make(chan string)

	// membuat fungsi yang akan membuat channel sebagai pengirim / mengirim data
	var sayHello = func(who string) {
		var data = fmt.Sprintf("Hello %s ", who)
		messageChannel1 <- data
	}

	// menjalankan fungsi sayHello() dengan goroutine
	go sayHello("Andi")
	go sayHello("Budi")
	go sayHello("Aan")

	// deklarasi variabel untuk menerima data dari channel message
	// pastikan untuk menerima data dari channel, ada data yang dikirim dari channel
	// dalam kasus ini ada 3 kali pengiriman data, maka hanya bisa menerima data 3 kali
	var message1 = <-messageChannel1
	fmt.Println(message1)

	var message2 = <-messageChannel1
	fmt.Println(message2)

	var message3 = <-messageChannel1
	fmt.Println(message3)

	// Channel sebagai tipe data parameter
	var messageChannel2 = make(chan string)

	// mengirim data lewat channel messageChannel2 dengan metode IIFE atau anonymous function langsung dipanggil atau isitilah nya lagi callback
	for _, each := range []string{"Andi", "Budi", "Sudarsono", "Aan"} {
		go func(who string) {
			var data = fmt.Sprintf("Hai %s ", who)
			messageChannel2 <- data
		}(each)
	}

	// menerima data lewat channel messageChannel2 dengan metode pengambilan dengan fungsi
	for i := 0; i < 4; i++ {
		printMessage(messageChannel2)
	}

}

func printMessage(what chan string) {
	fmt.Println(<-what)
}
