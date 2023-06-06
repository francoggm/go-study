package main

import "fmt"

func main() {
	var fruitList [4]string
	// fruitList := [4]string{}

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Banana"

	fmt.Println("List = ", fruitList)
	fmt.Println("Length = ", len(fruitList))

	vegList := [3]string{"potato", "beans", "mushroom"}

	fmt.Println("List = ", vegList)
}