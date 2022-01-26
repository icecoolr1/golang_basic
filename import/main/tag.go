package main

import (
	"fmt"
	"reflect"
)

type Resume struct {
	Name string `info:"name" doc:"我的名字"`
	sex string `info:"sex"`

}

func findTag(str interface{})  {
	t := reflect.TypeOf(str).Elem()
	for i := 0 ; i < t.NumField();i++{
		tagstring := t.Field(i).Tag.Get("info")
		fmt.Println("info:",tagstring)
	}
}

func main() {
	resume := Resume{
		Name: "wo",
		sex:  "1",
	}
	findTag(&resume)
}