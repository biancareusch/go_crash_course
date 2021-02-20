package main

import "fmt"

func main() {
	ids := []int{12, 13, 43, 21, 44}

	//Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	//Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	//Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	//Range with map
	emails := map[string]string{"Bob": "bob@gmail.com", "sharon": "sharon@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}
}