package main

import "fmt"

func main() {
	stu,_,_:=GetClass()
	fmt.Println(stu)
}
func GetClass()(stuNum int, className, headTeacher string) {
	return 9, "一班", "张三"
}

