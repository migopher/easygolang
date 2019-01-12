package main

import (
	"fmt"
)

/**
函数中的可变参数
在使用可变参数，可以传入一个切片但参数加...
 */
func main() {
	sum := arrAdd(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(sum)
	para := []int{1, 2, 3, 4, 5}
	fmt.Println(para)
	//fmt.Println(arrAdd(para)) //直接传入切片报错
	fmt.Println(arrAdd(para...)) //需要para...表示可变参数
}

/**
func arrAdd(a ...int) int {
}
可变参数常用于不确定的多个相同类型
比如传入不确定的整形进行计算

 */
func arrAdd(a ...int) int {
	var sum int
	for _, v := range a {
		sum += v
	}
	return sum
}
