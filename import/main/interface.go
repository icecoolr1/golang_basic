package main2

import "fmt"

// IStudent 本质是一个指针
type IStudent interface {
	Sleep()
	GetColor() string
	GetType() string
}

//具体的类
type Cat struct {
	color string
}


func (this *Cat) Sleep() {
	fmt.Println("cat is sleeping")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "cat"
}

func showAnimal(animal IStudent)  {
	animal.GetType()
	animal.Sleep()
}

func main() {
	var animal IStudent//接口数据类型 父类指针
	animal = &Cat{"Green"}
	animal.Sleep()

}
