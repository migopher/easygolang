package main

import "fmt"

type cloFun func() int
/**
函数的闭包
传入一个闭包函数
函数也是属于一种变量类型可以通过type进行定义
 */
func main()  {
	fmt.Println(varFun(func() int {
		return 100
	}))
}
func varFun(a cloFun) int{
	return a()
}