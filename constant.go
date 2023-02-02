package main

import "fmt"

func main() {
	const firstName string = "David"
	const lastName = "Geraldo"
	const age = 22

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(age)

	const (
		office = "Bank Hana"
		role   = "Software Engineer"
	)

	fmt.Println(office)
	fmt.Println(role)
}
