package main

import "fmt"

func main() {
	a := 5
	// assigning b to a pointer of a
	// b is the memory address of where a is stored
	b := &a

	fmt.Println(a, b)
	// This equals *int
	fmt.Printf("%T\n", b)

	// Use * to read val from address
	fmt.Println(*b)
	// This gives 5
	fmt.Println(*&a)

	// Change val with pointer
	*b = 10
	fmt.Println(a)
}