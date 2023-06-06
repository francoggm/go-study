package main

import "fmt"

func main() {
	// var ptr *int

	myNumber := 23

	var ptr = &myNumber

	fmt.Println("Pointer address =", ptr)
	fmt.Println("Pointer value =", *ptr)

	*ptr *= 5

	fmt.Println("Pointer value =", *ptr)
}