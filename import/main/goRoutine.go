package main

import (
	"runtime"
	"time"
)

func testGo(){
	println("方法A")
}


// main方法相当于一个线程（主协程） 线程之上有多个协程（从协程）
func main() {
	//i :=0
	for  {
		//i++
		//go testGo()
		println("Main方法")
		time.Sleep(1*time.Second)
		//匿名函数
		go func(){
			println("方法B")
			//单纯退出协程没啥用 因为线程还会不断开启协程
			time.Sleep(1*time.Second)
			//强制退出 整个go程
			runtime.Goexit()

		}()//加个()代表执行

		//退出线程会导致协程退出
		//return
	}



}
