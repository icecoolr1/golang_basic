package main

func main() {

	// 固定长度数组
	//var myArray1 [10] int
	//myArray2 := [3]int{1,2,3}

	//切片 == 动态数组 切片是引用传递(slice,map,channel) 数组是值传递
	myArray3 := []int{1, 2, 3}

	//使用range会返回两个值 第一个是索引 第二个是值
	//for i,i2 := range myArray1 {
	//	println(i)
	//	println(i2)
	//}
	for _, i2 := range myArray3 {
		println(i2)
	}

}
