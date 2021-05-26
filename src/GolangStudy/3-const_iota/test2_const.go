package main

import "fmt"

// const定义枚举
const (
	// 可以在const()添加一个关键字iota，每行的iota都会累加1，第一行的iota默认是0
	BEIJING = 10 * iota
	SHANGHAI
	SHENZHEN
)

const (
	a, b = iota + 1, iota + 2
	c, d
	e, f

	g, h = iota * 2, iota * 3
	i, k
)

func main() {
	// 常量(只读)
	const length int = 10

	fmt.Println("length =", length)

	fmt.Println("BEIJING =", BEIJING)
	fmt.Println("SHANGHAI =", SHANGHAI)
	fmt.Println("SHENZHEN =", SHENZHEN)

	fmt.Println("a,b：", a, b)
	fmt.Println("c,d：", c, d)
	fmt.Println("e,f：", e, f)
	fmt.Println("g,h：", g, h)
	fmt.Println("i,k：", i, k)
}
