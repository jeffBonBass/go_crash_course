package main

import "fmt"

func main() {
	// Define map
	//emails := make(map[string]string)

	// Assign kv
	//emails["Bob"] = "bob@gmail.com"
	//emails["Sharon"] = "sharon@gmail.com"
	//emails["Jeff"] = "jjj@gmail.com"

	// another way to declar and assign maps
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	fmt.Println(emails)
	fmt.Println(emails["Bob"])
	fmt.Println(len(emails))

	// Delete an entry
	delete(emails, "Bob")
	fmt.Println(emails)
}
