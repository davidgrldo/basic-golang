package main

import "fmt"

func main() {
	var name string

	name = "David"
	fmt.Println(name)

	name = "David Geraldo"
	fmt.Println(name)

	var job = "Software Engineer"
	fmt.Println(job)

	var age = 23
	fmt.Println(age)

	country := "Indonesia"
	fmt.Println(country)

	var (
		firstName = "David"
		lastName  = "Geraldo"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}
