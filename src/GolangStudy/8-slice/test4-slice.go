package main

import "fmt"

func main() {
	fmt.Println("------事先声明cap的情况------")

	var numbers = make([]int, 3, 5)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	// 向切片追加一个元素1
	numbers = append(numbers, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	// 向切片追加一个元素2
	numbers = append(numbers, 2)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	// 向切片追加一个元素3
	numbers = append(numbers, 3)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	fmt.Println("------没有事先声明cap的情况------")

	var numbers2 = make([]int, 3)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers2), cap(numbers2), numbers2)

	// 向切片追加一个元素
	numbers2 = append(numbers2, 3)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers2), cap(numbers2), numbers2)
}
