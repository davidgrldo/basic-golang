package main

import "fmt"

func main() {

	// Declare variable (array)
	var names [3]string

	names[0] = "David"
	names[1] = "Geraldo"
	names[2] = "Pakpahan"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// Declare variable (array) and set the value
	var values = [3]int{
		99, 98, 95,
	}

	fmt.Println(values)
	// Get data in index position
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	// Get the length of array
	fmt.Println(len(names))
	fmt.Println(len(values))

	var testArray [10]string

	fmt.Println(len(testArray))

	// Update data in index position
	values[0] = 100
	fmt.Println(values[0])
}
