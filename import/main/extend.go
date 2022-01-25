package main

import (
	"fmt"
)

type Human struct {
	name string
	sex string
}

type SuperHuman struct {
	Human//代表继承了human的方法
	level int
}

func (this *SuperHuman) Eat()  {
	fmt.Println("超人从不吃东西")
}


func (this *Human)Eat(){
	fmt.Println("吃")
}

func (this *Human)Drink(){
	fmt.Println("喝")
}

func main() {
	h:=Human{
		name: "wang",
		sex:  "very",
	}

	// 初始化
	//sp := SuperHuman{
	//	Human{"wang", "very"},
	//	10,
	//}

	//不初始化
	var sp SuperHuman

	//sp.name
	sp.Drink()


	h.Drink()
}
