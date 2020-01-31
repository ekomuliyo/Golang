package main

import "fmt"

func main() {

	var buah = []string{"apel", "jeruk", "nanas"} // ini slice
	var buah2 = [2]string{"jeruk", "jambu"}       // ini array
	fmt.Println(buah)
	fmt.Println(buah2)

	// perbedaan array dengan slice
	// -> array adalah kumpulan nilai atau elemen
	// -> slice referensi tiap elemen tsb, apbaila data referensi slice di ubah maka data awal nya juga diubahb

	fmt.Println()
	var buahLagi = []string{"apel", "jambu", "jeruk", "pempek"}

	var newBuah1 = buahLagi[0:2] // [apel jambu]
	var newBuah2 = buahLagi[0:4] // [apel jambu jeruk pempek]
	var newBuah3 = buahLagi[0:0] // []
	var newBuah4 = buahLagi[4:4] // []
	var newBuah5 = buahLagi[:]   // [apel jambu jeruk pempek]
	var newBuah6 = buahLagi[2:]  // [jeruk pempek]
	var newBuah7 = buahLagi[:2]  // [apel jambu]

	fmt.Println(newBuah1)
	fmt.Println(newBuah2)
	fmt.Println(newBuah3)
	fmt.Println(newBuah4)
	fmt.Println(newBuah5)
	fmt.Println(newBuah6)
	fmt.Println(newBuah7)

	fmt.Println()
	var buah2Lagi = []string{"pempek", "cuko", "model", "tekwan"}

	var buah2LagiA = buah2Lagi[0:3]
	var buah2LagiB = buah2Lagi[1:4]

	var buah2LagiAA = buah2LagiA[1:2]
	var buah2LagiBA = buah2LagiB[0:1]

	fmt.Println(buah2Lagi)
	fmt.Println(buah2LagiA)
	fmt.Println(buah2LagiB)
	fmt.Println(buah2LagiAA)
	fmt.Println(buah2LagiBA)

	fmt.Println()
	fmt.Println("pempek diubah jadi gorengan")

	buah2LagiBA[0] = "kecap"
	fmt.Println(buah2LagiA)
	fmt.Println(buah2LagiB)
	fmt.Println(buah2LagiAA)
	fmt.Println(buah2LagiBA)

	fmt.Println()
	// fungsi len(slice) / length
	fmt.Println("jumlah elemen =", len(buah2Lagi))

	fmt.Println()
	// fungsi cap(sice) / capacity
	var mobil = []string{"honda jazz", "hrv", "mobilio", "brio"}
	fmt.Println(len(mobil)) // 4
	fmt.Println(cap(mobil)) // 4

	fmt.Println()
	var aMobil = mobil[0:3]
	fmt.Println(len(aMobil)) // 3
	fmt.Println(cap(aMobil)) // 4

	fmt.Println()
	var bMobil = mobil[1:4]
	fmt.Println(len(bMobil)) // 3
	fmt.Println(cap(bMobil)) // 3

	fmt.Println()
	fmt.Println()
	// fungsi append(slice, "object slice baru")
	// apabila len == cap maka nilai baru akan ditambahkan ditempatkan sebagai referensi baru
	// apabila len < cap maka nilai yang baru ditambahkan ditempatkan kedalam cakupan kapasitas,
	// menjadikan slice lain yang menjadi referensi nya akan berubah juga

	var motor = []string{"jupiter mx", "jupiter z", "v-ixion", "mio"}
	fmt.Println(motor)

	var aMotor = append(motor, "vespa")
	fmt.Println(aMotor)

	var bMotor = motor[0:3]
	fmt.Println(len(bMotor))
	fmt.Println(cap(bMotor))

	fmt.Println(motor)
	fmt.Println(bMotor)

	fmt.Println()
	var cMotor = append(bMotor, "beat")
	fmt.Println(motor)
	fmt.Println(bMotor)
	fmt.Println(cMotor)

	fmt.Println()
	// fungsi copy(destination, source)
	// fungsi copy mengembalikan nilai jumlah elemen yang berhasil copy
	destination := make([]string, 3)
	source := []string{"apel", "jeruk", "semangka", "rambutan"}
	result := copy(destination, source)

	fmt.Println(destination)
	fmt.Println(source)
	fmt.Println(result)

	fmt.Println()
	dst := []string{"pempek", "pempek", "pempek"}
	src := []string{"bakso", "tekwan"}
	result2 := copy(dst, src)

	fmt.Println(dst)
	fmt.Println(src)
	fmt.Println(result2)

	fmt.Println()
	// mengakses elemen slice dengan 3 index
	var laptop = []string{"asus", "acer", "hp", "lenovo"}
	var aLaptop = laptop[0:2]
	var bLaptop = laptop[0:2:2]

	fmt.Println(laptop)
	fmt.Println(len(laptop))
	fmt.Println(cap(laptop))

	fmt.Println(aLaptop)
	fmt.Println(len(aLaptop))
	fmt.Println(cap(aLaptop))

	fmt.Println(bLaptop)
	fmt.Println(len(bLaptop))
	fmt.Println(cap(bLaptop))
}
