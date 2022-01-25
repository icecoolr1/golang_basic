package main

type Hero struct {
	//变量和方法首字母大写为访问 首字母小写为私有
	name string
	ad int
	level int
}


// GetName Hero不带*默认值传递
func (this *Hero)GetName() string {
	return this.name
}

func (this *Hero)SetName(newName string)  {
	this.name = newName
}

func main() {
	hero:= Hero{name: "zzz",ad: 10,level: 20}
	hero.SetName("haha")
	println(hero.GetName())
}

