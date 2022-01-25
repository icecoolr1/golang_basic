package main

import "fmt"

func main() {
	// 声明slice1是一个切片 并且初始化 推荐
	slice1 := []int{1, 23, 3}
	//len=3 slice=[1 23 3]
	fmt.Printf("len=%d slice=%v\n", len(slice1), slice1)

	//声明是一个切片 但是没有分配空间
	//var slice2 []int
	//开辟空间
	//slice2 = make([]int,3)

	//等价
	//var slice3 []int = make([]int,3)


	//推荐
	//slice3 := make([]int,1)

	//容量追加

	// 5代表容量 默认为len的值
	var numbers = make([]int, 3, 5)
	println(cap(numbers))
	//像numbers切片追加元素1
	numbers = append(numbers, 1)

	//区间截取 左闭右开 指向同一数组
	s := []int{1, 2, 3}
	s1 := s[0:2]
	for _, i := range s1 {
		println(i)
	}
	//深拷贝
	s2 := make([]int, 3)
	//s->s2
	copy(s2, s)

}
