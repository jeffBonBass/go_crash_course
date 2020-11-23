package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for uint32
	// float32 float64
	// complex64 complex128

	// using var
	//var name = "Jeff"
	//name := "Jeff" //shorthand var declaration
	//email := "jeff@gmail.com"

	name, email := "Jeff", "jeffb@gmall.com"
	var age = 57
	const isCool = true

	size := 1.3

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", size)
}
