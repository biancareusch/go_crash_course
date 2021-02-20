package main

import "fmt"

func main() {
	//Arrays
	var fruitArr [2]string

	//Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	fmt.Println(fruitArr)

	//Declare and assign
	nameArr := [2]string{"Bob", "Mike"}

	fmt.Println(nameArr)

	//Slices
	fruitSlice := []string{"Apple", "Orange", "Grape", "Cherry"}

	fmt.Println(len(fruitSlice)) //length
	fmt.Println(fruitSlice[1:3]) //range
}
