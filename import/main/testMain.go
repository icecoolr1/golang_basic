package main

import (

	// 匿名 import
	//_ "testImport/lib1"

	// 别名
	//lib "testImport/lib1"

	//  直接配到lib1下
	. "testImport/lib1"
)

func main() {
	TestLib1()
}
