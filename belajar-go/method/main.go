package main

import (
	"fmt"
	"strings"
)

func main() {
	var s1 = student{"andi pratama", 2}
	s1.sayHello()

	var name = s1.getName(2)
	fmt.Println("nama panggilan :", name)

	fmt.Println()

	// method pointer

	var s2 = student{"rudi saputra", 4}
	fmt.Println("nilai awal (s2) 		:", s2.name)

	s2.changeName("ubah rudi") // hasil nya tidak berubah, yang berubah nilainya cuman di dalam method nya saja
	fmt.Println("nilai setelah di ubah (s2) 	:", s2.name)

	s2.changeNamePointer("ubah rudi lagi") // hasilnya berubah
	fmt.Println("nilai setelah ubah (s2) 	:", s2.name)

}

type student struct {
	name  string
	grade int
}

func (s student) sayHello() {
	fmt.Println("halo", s.name)
}

func (s student) getName(index int) string {
	return strings.Split(s.name, " ")[index-1]
}

func (s student) changeName(name string) {
	fmt.Println("changeName 			:", name)
	s.name = name
}

func (s *student) changeNamePointer(name string) {
	fmt.Println("changeNamePointer 		:", name)
	s.name = name
}
