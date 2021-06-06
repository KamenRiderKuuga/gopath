package main

import "fmt"

func printArray(myArray [4]int) {
	// 值拷贝
	for index, value := range myArray {
		fmt.Println("index =", index, "value =", value)
	}

	myArray[0] = 111
}

func main() {
	// 固定长度的数组
	var myArray1 [10]int

	myArray2 := [10]int{1, 2, 3, 4}
	myArray3 := [4]int{11, 22, 33, 44}

	for i := 0; i < len(myArray1); i++ {
		fmt.Println(myArray1[i])
	}

	for index, value := range myArray2 {
		fmt.Println("index =", index, "value =", value)
	}

	// 查看数组的数组类型
	fmt.Printf("type of myArray1 is %T\n", myArray1)
	fmt.Printf("type of myArray2 is %T\n", myArray2)
	fmt.Printf("type of myArray3 is %T\n", myArray3)

	// 这里因为指定了数组类型，只能传递myArray3
	printArray(myArray3)
	fmt.Println("----------")
	for index, value := range myArray3 {
		fmt.Println("index =", index, "value =", value)
	}
}
