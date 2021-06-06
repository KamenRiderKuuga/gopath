package main

import "fmt"

func main() {
	// 声明一个切片，并以初始值1，2，3初始化，长度len为3
	slice1 := []int{1, 2, 3}

	fmt.Printf("len = %d, slice1 = %v\n", len(slice1), slice1)

	// 声明一个切片，但是没有为其分配空间
	var slice2 []int
	// 为其分配3个空间，初始化值是0
	slice2 = make([]int, 3)

	// 简写
	// var slice2 = make([]int, 3)
	// slice2 := make([]int, 3)

	fmt.Printf("len = %d, slice2 = %v\n", len(slice2), slice2)

	var slice3 []int
	if slice3 == nil {
		fmt.Println("slice3是一个空切片")
	} else {
		fmt.Println("slice3是有空间的")
	}
}
