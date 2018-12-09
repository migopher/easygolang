package main

import "fmt"

var value1 int //全局声明变量,可不使用


func main()  {
	var value3 int32
	value3=32
	value1=32
	value2:=64 //函数内声明变量,必须使用
	fmt.Println(value2)
	value1,value2=value2,value1 //互换值
	fmt.Println(value1,value2)
	if value2 == int(value3) {  //不同类型无法进行比较 使用int()强制转换
		fmt.Println("相同变量类型")
	}
}