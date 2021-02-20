package main

import "fmt"

func main() {
	//Define map
	// emails := make(map[string]string)

	// // Assign key value
	// emails["Bob"] = "bob@gmail.com"
	// emails["Sharon"] = "sharon@gmail.com"
	// emails["Mike"] = "mike@gmail.com"

	// Declare map and add keyvalue
	emails := map[string]string{"Bob": "bob@gmail.com", "sharon": "sharon@gmail.com"}

	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])
	fmt.Println(emails)

	//Delete from map
	delete(emails, "bob")
	fmt.Println(emails)
}
