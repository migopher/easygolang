package main

import (
	"fmt"
	"reflect"
)

func main() {
	//建议浮点型float 使用float64省下转换工作
	//math包函数接收要求都为float64
	var value1 float32
	value1=3.14
	value2 := 3.14 //默认浮点型为float64,有小数点推倒为浮点型
	fmt.Println(reflect.TypeOf(value2)) //反射出变量类型
	if float64(value1) == value2 { //强制转换成float64,浮点精准度不同会导致比较结果与预期不一样
		fmt.Println("值相同")
	}else {
		fmt.Println("值不相同")
	}
}
