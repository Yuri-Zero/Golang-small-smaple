package myjson

type MyJsonType struct {
	Students []Student `json:"student"`
}
type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Class Class  `json:"class"`
}
type Class struct {
	Grade10 string `json:"grade10"`
	Grade9  string `json:"grade9"`
	Grade8  string `json:"grade8"`
}
