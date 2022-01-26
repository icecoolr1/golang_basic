package main

import (
	"fmt"
	"reflect"
)

func reflectNUm(arg interface{})  {
	fmt.Println(reflect.TypeOf(arg))
	fmt.Println(reflect.ValueOf(arg))
}


func main() {
	var num float32 = 1.234
	reflectNUm(num)
}
