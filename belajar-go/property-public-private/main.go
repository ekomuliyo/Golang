package main

import (
	// "Golang/belajar-go/property-public-private/library"
	"Golang/belajar-go/property-public-private/library"
	. "Golang/belajar-go/property-public-private/library"
	f "fmt"
)

func main() {
	library.SayHello()
	// library.introduce("eko")

	f.Println()

	var student1 = library.Student{"andi", 12}
	f.Println("Name ", student1.Name)
	f.Println("grade ", student1.Grade)

	f.Println()

	var student2 = Student{"budi", 13}
	f.Println("name ", student2.Name)
	f.Println("grade ", student2.Grade)

	f.Println()
	sayHello()

	f.Println()
	f.Printf("Name : %s\n", Student2.Name)
	f.Printf("Grade : %d\n", Student2.Grade)
}
