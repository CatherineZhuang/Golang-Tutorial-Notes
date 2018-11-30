package main

import "fmt"

func main() {
	// Using var
	var name string = "Brad"
	fmt.Println(name)
	// if you declared a variable in go, you have to use it
	var age int32 = 37
	fmt.Println(age)
	fmt.Printf("%T\n", age)

	// Using const
	const isCool = true
	// isCool = false ERRORRR
	fmt.Println(isCool)

	// Shorthand declaration
	hello := "Hello World"
	fmt.Println(hello)

	username, email := "Catherine", "catherine@gmail.com"
	fmt.Println(username, email)
}

// MAIN TYPES
// string
// bool
// int
// int int8 int16 int32 int 64
// uint uint8 uint16 uint32 uint64 uintptr
// byte - alias for uint8
// rune - alias for int32
// float32 float64
// complex64 complex128