package main

import "fmt"

func main() {

	// penerapan struct sederhana
	var s1 student
	s1.name = "andi pratama"
	s1.grade = 3

	fmt.Println("name :", s1.name)
	fmt.Println("grade :", s1.grade)

	fmt.Println()

	// inisialisasi object struct
	var s2 = student{}
	s2.name = "budi"
	s2.grade = 1

	// var s3 = student{"eko", 3}

	var s4 = student{name: "rudi", grade: 4}

	fmt.Println("student 2 :", s2.name)
	// fmt.Println("student 3 :", s3.name)
	fmt.Println("student 4 :", s4.name)

	fmt.Println()

	// variabel object pointer
	var s5 = student{name: "yudi", grade: 5}

	var s6 *student = &s5

	fmt.Println("student 5  :", s5.name)
	fmt.Println("student s6 :", s6.name)

	s6.name = "yudi pratama"
	fmt.Println("student 5  :", s5.name)
	fmt.Println("student s6 :", s6.name)

	fmt.Println()

	// Embedded Struct

	var s7 = student{}
	s7.name = "riko"
	s7.age = 21
	s7.grade = 2

	fmt.Println("name :", s7.name)
	fmt.Println("age :", s7.age)
	fmt.Println("age person :", s7.person.age)
	fmt.Println("grade :", s7.grade)

	fmt.Println()

	// Embedded Struct dengan property yang sama
	var s8 = student{}
	s8.name = "aan"
	s8.age = 4        // age of student
	s8.person.age = 3 // age of person

	fmt.Println("name :", s8.name)
	fmt.Println("age :", s8.age)
	fmt.Println("age person :", s8.person.age)

	fmt.Println()

	// Pengisian nilai sub-struct
	var p1 = person{name: "gio", age: 17}
	var s9 = student{grade: 6, person: p1, age: 19}

	fmt.Println("name  :", s9.person.name)
	fmt.Println("age   :", s9.age)
	fmt.Println("grade :", s9.grade)

	fmt.Println()

	// Anonymous Struc, cocok untuk menggunakan struct sekali pakai

	// tanpa pengisian property
	var s10 = struct {
		person
		grade int
	}{}

	var s11 = struct {
		person
		grade int
	}{
		person: person{"jui", 25},
		grade:  6,
	}

	s10.person = person{"andi", 22}
	s10.grade = 3

	fmt.Println("name  :", s10.person.name)
	fmt.Println("age   :", s10.person.age)
	fmt.Println("grade :", s10.grade)
	fmt.Println()
	fmt.Println("name  :", s11.person.name)
	fmt.Println("age   :", s11.person.age)
	fmt.Println("grade :", s11.grade)

	fmt.Println()

	// kombinasi slice & struct
	var allPersons = []person{
		{name: "budi", age: 15},
		{name: "aan", age: 17},
	}

	for _, person := range allPersons {
		fmt.Println("name =", person.name)
		fmt.Println("age =", person.age)
	}

	fmt.Println()

	// Inisialisasi slice anonymous struct
	var allStudents = []struct {
		person
		grade int
	}{
		{person: person{"budi", 21}, grade: 4},
		{person: person{"aan", 20}, grade: 3},
		{person: person{"riko", 19}, grade: 2},
	}

	for _, student := range allStudents {
		fmt.Println(student)
	}

	fmt.Println()

	// Deklarasi anonymous struct menggunakan keyword var
	student2.grade = 10
	fmt.Println(student2.grade)

	// deklarasi sekaligus inisialisasi
	var student3 = struct {
		grade int
	}{
		12,
	}
	fmt.Println(student3.grade)

	fmt.Println()

	// Type Alias
	type Person struct {
		name string
		age  int
	}

	type People = Person

	var p2 = Person{"andi", 20}
	fmt.Println(p2)

	var p3 = People{"budi", 21}
	fmt.Println(p3)

	// casting lewat alias
	people := People{"aan", 22}
	fmt.Println(Person(people))

	person := Person{"gio", 23}
	fmt.Println(Person(person))
}

type student struct {
	name  string
	grade int
	person
	age int
}

var student2 struct {
	grade int
}

type person struct {
	name string
	age  int
}

// nested struct
type student3 struct {
	person struct {
		name string
		age  int
	}
	grade   int
	hobbies []string
}

// tag property dalam struct
type person2 struct {
	name string `tag1`
	age  int    `tag2`
}

// pembuatan struc dengan teknik alias
type person4 = struct {
	name string
	age  int
}
