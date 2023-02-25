package main

import (
	"fmt"
	"jsonparser/internal/myjson"
)

func main() {
	filepath := "../assets/test.json"
	val := myjson.JsonFileReder(filepath)
	stus := myjson.Deserialize(val)
	for i := 0; i < len(stus.Students); i++ {
		fmt.Println("姓名:", stus.Students[i].Name)
		fmt.Println("年龄:", stus.Students[i].Age)
		fmt.Println("高一班级:", stus.Students[i].Class.Grade9)
		fmt.Println("高二班级:", stus.Students[i].Class.Grade8)
		fmt.Println("高三班级:", stus.Students[i].Class.Grade10)
	}
	myjson.Append()
}
