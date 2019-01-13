package main

import "fmt"

/**
闭包函数
匿名函数同样被称为闭包函数
 */
func main() {
	/**
	Add不需参数执行返回闭包函数，在闭包函数中传入一个整数进行计算
	 */
	fmt.Println(Add()(10))
}

/**
Add函数返回一个闭包函数，返回出去直接执行
可使用于某个函数捕捉一些外部异常状态
 */
func Add() func(b int) int {
	return func(b int) int {
		return b * 10
	}
}
