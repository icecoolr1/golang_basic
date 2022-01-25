package main

import "fmt"

/*
	和C++差不多
*/

func swap(a *int, b *int) {
	temp := 0
	temp = *a
	*a = *b
	*b = temp
}

func main() {
	a := 3
	b := 4
	swap(&a, &b)
	fmt.Println("a=", a, "b=", b)
	// 相当于析构函数 多个defer看成栈结构 后进先出
	defer println("执行.....")

}
