package main

import "fmt"

// 声明全局变量 方法一、方法二、方法三是可以的
// 使用方法四则是不可以的，其只能在函数体内进行声明
var gA int = 100
var gB = 200

func main() {
	// 1. 声明变量，使用其默认值
	var a int
	fmt.Println("a  = ", a)
	fmt.Printf("type of a = %T\n", a)

	// 2. 声明一个变量，初始化一个值
	var b int = 100
	fmt.Println("b =", b)
	fmt.Printf("type of b = %T\n", b)

	var bb string = "abcd"
	fmt.Printf("bb == %s,type of bb = %T\n", bb, bb)

	// 3. 通过值自动匹配当前的数据类型
	var c = 100
	fmt.Println("c = ", c)
	fmt.Printf("type of b = %T\n", c)

	// 4. 常用的方法，既初始化，又赋值
	e := 100
	fmt.Println("e =", e)
	fmt.Printf("type of e = %T\n", e)

	// 声明多个变量
	var xx, yy int = 100, 200
	fmt.Println("xx =", xx, "yy =", yy)

	var kk, ll = 100, "HANABI"
	fmt.Println("kk =", kk, "ll =", ll)

	// 多行的多变量声明
	var (
		vv int  = 100
		jj bool = true
	)

	fmt.Println("vv =", vv, "jj =", jj)
}
