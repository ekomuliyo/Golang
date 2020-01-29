package main

import "fmt"

func main() {

	// menggunakan if, else if, else
	var angka = 10

	if angka == 10 {
		fmt.Println("Lulus dengan sempurna")
	} else if angka > 5 {
		fmt.Println("Lulus aja")
	} else if angka == 4 {
		fmt.Println("Dikit lagi lulus")
	} else {
		fmt.Printf("Tidak lulus, silahkan coba lagi")
	}

	// menggunakan variabel temporary (sementara) pada if, else
	var angka2 = 9000

	if persen := angka2 / 100; persen >= 100 { // nilai dari variabel persen adalah hasil perhitungannya dan nilai hasil dari pembagian akan terus dibandingkan ke kondisi selanjunya
		fmt.Printf("%d %s sangat memuaskan!\n", persen, "%")
	} else if angka2 >= 70 {
		fmt.Printf("%d %s memuaskan\n", persen, "%")
	} else {
		fmt.Printf("%d %s cukup\n", persen, "%")
	}

	// menggunakan switch case
	var angka3 = 10

	switch angka3 {
	case 10:
		fmt.Println("Sangat memuaskan")
	case 7:
		fmt.Println("cukup memuaskan")
	default:
		fmt.Println("lumayanlah")
	}

	// mengunakan switch case dengan banyak kondisi
	var angka4 = 8

	switch angka4 {
	case 8:
		fmt.Println("Keren")
	case 7, 6, 5, 4:
		fmt.Println("cukup keren")
	default:
		fmt.Println("cukuplah")
	}

	// menggunakan switch case dengan kurung kurawal
	var angka5 = 7

	switch angka5 {
	case 8:
		fmt.Println("sangat memuaskan")
	case 7, 6, 5, 4:
		fmt.Println("cukup memuaskan")
	default:
		{
			fmt.Println("cukuplah")
			fmt.Println("masih bisa dipertimbangkan")
		}
	}

	// menggunakan switch dengan gaya if else
	var angka6 = 7

	switch {
	case angka6 == 8:
		fmt.Println("Sangat memuaskan")
	case (angka6 < 8) && (angka6 > 3):
		fmt.Println("cukup memuaskan")
	default:
		{
			fmt.Println("cukuplah")
			fmt.Println("masih bisa dipertimbangkan")
		}
	}

	// menggunakan fallthrough pada switch case
	var angka7 = 7
	switch {
	case angka7 == 8:
		fmt.Println("sangat memuaskan")
	case (angka7 < 8) && (angka7 >= 5):
		fmt.Println("cukup memuaskan")
		fallthrough
	case angka7 < 5:
		fmt.Println("masih perlu belajar lagi")
	default:
		{
			fmt.Println("cukuplah")
			fmt.Println("masih bisa dipertimbangkan")
		}
	}

	// menggunakan kondisi bersarang
	var angka8 = 10

	if angka8 > 7 {
		switch angka8 {
		case 10:
			fmt.Println("sangat memuaskan")
		default:
			fmt.Println("bagus")
		}
	} else {
		if angka8 == 5 {
			fmt.Println("cukup")
		} else if angka8 == 3 {
			fmt.Println("coba lagi")
		} else {
			fmt.Println("kamu bisa kok")
			if angka8 == 0 {
				fmt.Println("kamu harus mencoba lagi!")
			}
		}
	}
}
