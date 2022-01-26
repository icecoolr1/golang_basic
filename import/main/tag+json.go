package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title string `json:"title"`
	Year int `json:"year"`
	Price float32 `json:"price"`
}

func main() {
	movie :=Movie{
		Title: "wwww",
		Year:  0,
		Price: 0,
	}
	//编码转json
	jsonStr,err:=json.Marshal(movie)
	if err!= nil{
		fmt.Println("导出json失败")
		return
	}

	fmt.Printf(string(jsonStr))


	//解码过程
	jsonMovie := Movie{}
	err = json.Unmarshal(jsonStr,&jsonMovie)
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println(jsonMovie)

}