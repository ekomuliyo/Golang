package main

import (
	"fmt"
	"reflect"
)

func main() {

	var number = 23

	var reflectValue = reflect.ValueOf(number)

	fmt.Println("type variable :", reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("value variable :", reflectValue.Int())
	}

	fmt.Println()
	// interface
	fmt.Println("nilai variabel :", reflectValue.Interface()) // ini hanya menampilkan string dari nilai aslinya apabila ingin mengoperasikannya maka lakukan casting sesuai keinginan reflectValue.Interface.(int)

	fmt.Println()
	// mengakses informasi properti variabel object
	var s1 = &student{Name: "andi", Grade: 10}
	s1.getPropertyInfo()

	fmt.Println()
	// mengakses informasi method variabel object
	var s2 = &student{Name: "budi", Grade: 12}
	fmt.Println("nama :", s2.Name)

	var reflectValue3 = reflect.ValueOf(s2)
	var method = reflectValue3.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("wick"),
	})

	fmt.Println("nama	:", s2.Name)
}

type student struct {
	Name  string
	Grade int
}

func (s *student) getPropertyInfo() {
	var reflectValue2 = reflect.ValueOf(s)

	if reflectValue2.Kind() == reflect.Ptr {
		reflectValue2 = reflectValue2.Elem()
	}

	var reflectType = reflectValue2.Type()

	for i := 0; i < reflectValue2.NumField(); i++ {
		fmt.Println("nama variabel	:", reflectType.Field(i).Name)
		fmt.Println("tipe data variabel:", reflectType.Field(i).Type)
		fmt.Println("nilai variabel:", reflectValue2.Field(i).Interface())
		fmt.Println("")
	}
}

func (s *student) SetName(name string) {
	s.Name = name
}
