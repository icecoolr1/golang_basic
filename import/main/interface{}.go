package main

import "fmt"

func myFunc(arg interface{})  {
	fmt.Println("myFunc is called")
	fmt.Println(arg)

	_,ok:=arg.(string)
	if !ok{
		println("不是String")
	}else {
		println("是String")
	}

}

type User struct {
	auth string
}

func main() {
	book := User{auth: "wo"}
	myFunc(book)
}