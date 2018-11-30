package main

import "fmt"

func main() {
	// Arrays
	var fruitArr [2]string

	// Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	fmt.Println(fruitArr)

	// Declare and assign at the same time
	fruitArr2 := [2]string{"Apple", "Orange"}
	fmt.Println(fruitArr2)

	// Slice doesnt have to be fixed length
	fruitSlice := []string{"Apple", "Orange", "Grape", "Banana"}
	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:2])
}