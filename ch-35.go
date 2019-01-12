package main

import (
	"fmt"
)

/**
匿名函数
不指定函数名的一种一种函数方式
 */
func main() {
	a := 21
	b := 10
	c := func(a, b int) int { return a * b }
	fmt.Println(c(a, b))
	/**
	这种也是可以使用
	 */
	func(a, b int) int {
		return a + b
	}(10, 10) //从结尾处传入参数直接调用函数,最后一对括号表示对该匿名函数的调用
}
