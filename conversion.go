package main

import "fmt"

func main() {
	var nilai32 = 129
	var nilai64 = int64(nilai32)
	var nilai8 = int8(nilai64)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var name = "David"
	var getString byte = name[0]
	var convertedString = string(getString)

	fmt.Println(name)
	fmt.Println(getString)
	fmt.Println(convertedString)
}
