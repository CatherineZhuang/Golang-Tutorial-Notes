package main

import "fmt"

func main() {
	// Map is key value pairs

	// Define map
	emails := make(map[string]string)
	// Assign key values
	emails["Bob"] = "bob@gmail.com"
	emails["Alex"] = "alex@gmail.com"
	emails["Mike"] = "mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["Bob"])

	// Delete from map
	delete(emails, "Bob")
	fmt.Println(emails)

	// Delcare map and add kv
	email := map[string]string{"Bob":"bob@gmail.com", "Alex":"alex@gmail.com"}
	fmt.Println(email)
}