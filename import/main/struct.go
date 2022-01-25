package main

// 声明int的别名myint
type myint int

// Book 定义一个结构体
type Book struct {
	title string
	author string
}


func main() {
	var book Book
	book.title="111"
	book.author="222"
}


