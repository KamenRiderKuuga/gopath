package main

import "fmt"

func printArray(myArray []int) {
	// 引用传递
	for _, value := range myArray {
		fmt.Println("value = ", value)
	}

	myArray[0] = 100
}

func main() {
	myArray := []int{1, 2, 3, 4}

	fmt.Printf("type of myArray %T\n", myArray)

	printArray(myArray)

	fmt.Println("----------")

	for _, value := range myArray {
		fmt.Println("value = ", value)
	}
}
