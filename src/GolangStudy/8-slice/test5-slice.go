package main

import "fmt"

func main() {
	s := []int{1, 2, 3} // len = 3, cap = 3

	// [0:2)
	s1 := s[0:2] // [1,2]

	fmt.Println(s1)

	s1[0] = 100

	fmt.Println("s的内容是", s, "s1的内容是", s1)
	fmt.Printf("s1的len = %d, cap = %d, value = %v\n", len(s1), cap(s1), s1)

	s1 = append(s1, 200)
	fmt.Println("append之后，s的内容是", s, "s1的内容是", s1)
	fmt.Printf("append之后，s1的len = %d, cap = %d, value = %v\n", len(s1), cap(s1), s1)

	s1 = append(s1, 300)
	fmt.Println("再次append之后，s的内容是", s, "s1的内容是", s1)
	fmt.Printf("再次append之后，s1的len = %d, cap = %d, value = %v\n", len(s1), cap(s1), s1)

	s2 := make([]int, 3) // s2 = [0,0,0]

	// 将s中的值，依次拷贝到s2中
	copy(s2, s)
	s2[0] = 0
	fmt.Printf("拷贝并调整s2成员之后，s的值为%v,s2的值为%v\n", s, s2)

	// 尝试拷贝一个二维数组
	s3 := make([][]int, 3)
	s3[0] = make([]int, 2)
	s3[1] = make([]int, 2)
	s3[2] = make([]int, 2)
	s4 := make([][]int, 3)
	copy(s4, s3)
	s3[0][0] = 1000
	fmt.Printf("拷贝并调整s3成员之后，s3的值为%v,s4的值为%v\n", s3, s4)
}
