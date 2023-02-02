package main

import "fmt"

func main() {
	type KTP string
	type status bool

	var noKTP KTP = "1234234534564567"
	var isMarried status = true
	fmt.Println(noKTP)
	fmt.Println(isMarried)
}
