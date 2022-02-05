package main

import (
	"fmt"
)

/*
	变量声明方式
	声明全局变量 1,2可以 3只能用在函数体内

 */


//别名
type myint int


func main() {
	//声明一个变量 默认的值是0
	// 1. var a int = 8
	//2 .
	var a = 8
	fmt.Println(a)
	//3.既声明又赋值
	b := 9
	println(b)

	//声明多个变量
	//var xx,yy int = 100  ,200
	//var kk,ll = 100,"ABCD"

	// 多行多变量类型
	//var(
	//	vv int = 100
	//	jj bool = true
	//)
	//

	//常量
	const age int = 10
	fmt.Println(age)

	//const来定义枚举类型 可以添加关键字 iota 每行都会加1
	const(
		BeiJing = iota // iota = 0
		ShangHai // 其余可以不写  为1
		NanJing  // 为2
	)
}
