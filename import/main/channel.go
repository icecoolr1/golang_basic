package main

import (
	"fmt"
	"time"
)

func main() {
	// 定义一个单缓冲channel 数据类型为int
	//c:=make(chan int)
	//多缓冲
	cs:=make(chan int,3)
	go func() {
		defer println("结束")
		println("go运行")
		for i :=0 ; i < cap(cs) ; i++{
			cs <- i
		}
		//关闭通道
		close(cs)
	}()

	time.Sleep(2*time.Second)
	//从channel C 中读数据
	//num := <-cs

	//for i :=0 ; i < cap(cs) ; i++{
	//	num := <- cs
	//	fmt.Println(num)
	//}

	for{
		//ok为true channel没有关闭
		if data,ok := <-cs; ok{
			fmt.Println(data)
		}else {
			break
		}
	}


	//读数据但抛弃
	//<-c
	//ok代表是否成功
	//value,ok:=<-c
	//fmt.Println(num)




}
