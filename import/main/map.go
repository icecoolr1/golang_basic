package main

func main() {
	//声明map
	var myMap1 map[string]string
	//给map分配空间
	myMap1 = make(map[string]string, 10)
	//添加map
	myMap1["1"] = "one"
	myMap1["2"] = "two"

	for key,value := range myMap1 {
		println(key,value)
	}

	//删除
	delete(myMap1,"1")

	//可以不写length slice要写
	//myMap2 := make(map[string]string)

	// 初始化则不需要make
	//myMap3 := map[string] string{
	//	"1" : "one",
	//}

}
