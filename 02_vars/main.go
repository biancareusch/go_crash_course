package main

import "fmt"

var name = "Bianca"

func main() {
	//MAIN TYPES
	//string
	//bool
	//int
	//int
	//int int8 int16 int32 int64
	//uint uint8 uint16 uint32 uint32 uintptr
	//byte - alias for uint8
	//rune - alias for int32
	//float32 float64
	//complex64 complex128

	lastName := "Reusch" //declare variable without var, const, only available inside function
	var age int32 = 27   //casts datatype
	const isCool = true

	name, email := "Bianca", "bebe@gmail.com"
	size := 1.3

	fmt.Println(name, age, lastName, isCool, email)
	fmt.Printf("%T\n", size)
}
