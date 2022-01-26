package main2

import "fmt"

func myFunc(arg interface{})  {
	fmt.Println("myFunc is called")
	fmt.Println(arg)
}

type User struct {
	auth string
}

//func main() {
//	book := User{auth: "wo"}
//	myFunc(book)
//}