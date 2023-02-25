package myjson

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func JsonFileReder(filename string) []byte {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Open Json File Successfully")
	value, _ := io.ReadAll(jsonFile)
	defer jsonFile.Close()
	return value
}
func Deserialize(values []byte) MyJsonType {
	var students MyJsonType
	json.Unmarshal(values, &students)
	return students
}
func Append() {
	Slice := make([]Student, 0)
	Class := Class{"高三10班", "高二9班", "高一8班"}
	newstu := Student{
		Name:  "周润发",
		Age:   18,
		Class: Class,
	}
	newstu2 := Student{
		Name:  "成龙",
		Age:   18,
		Class: Class,
	}
	Slice = append(Slice, newstu)
	Slice = append(Slice, newstu2)
	var newJsonData MyJsonType = MyJsonType{Slice}
	Jsonstr, _ := json.Marshal(newJsonData)
	fmt.Println(string(Jsonstr))
}
