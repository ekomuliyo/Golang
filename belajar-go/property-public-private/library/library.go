package library

import "fmt"

// SayHello : method public untuk menampilkan pesan hello
func SayHello() {
	fmt.Println("hello")
	introduce("eko")
}

func introduce(name string) {
	fmt.Println("halo saya ", name)
}

type Student struct {
	Name  string
	Grade int
}

var Student2 = struct {
	Name  string
	Grade int
}{}

// fungsi init pertama kali dijalankan sebelum main
func init() {
	Student2.Name = "aan saputra"
	Student2.Grade = 2
	fmt.Println("--> run library/library.go")
}
