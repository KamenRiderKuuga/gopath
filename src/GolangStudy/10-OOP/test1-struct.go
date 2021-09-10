package main

import "fmt"

// 使用别名声明数据类型
type myint int

// 定义一个结构体
type Book struct {
	title string
	auth  string
}

func changeBook(book Book) {
	// 传递一个book的副本
	book.auth = "666"
}

func changeBook2(book *Book) {
	// 传递指针
	book.auth = "111"
}

func main() {
	var book1 Book
	book1.title = "Golang"
	book1.auth = "zhang3"

	fmt.Printf("%v\n", book1)

	changeBook(book1)

	fmt.Printf("%v\n", book1)

	changeBook2(&book1)

	fmt.Printf("%v\n", book1)
}
