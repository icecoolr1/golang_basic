package main

import (
	"fmt"
)

// 返回值放在后面
func bool(a string , b int) int{
	fmt.Println("a = ",a)
	fmt.Println("b = ",b)
	return 10
}

func foo2(a string,b int)(int,int)  {
	return 666,777
}

//带形参名称的返回值
func foo3(a string,b int)(r1 int,r2 int)  {
	r1 = 100
	r2 = 500
	return
}

func main() {
	// go语言的变量声明应该要让编译器知道它的数据类型
	c := bool("abc",555)
	fmt.Println(c)
	ret1, ret2 := foo2("haha",999)
	fmt.Println(ret1)
	println(ret2)
}
