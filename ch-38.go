package main

import "fmt"

/**
递归函数
递归函数可以很优雅的解决很多问题
比如:导航递归，栏目递归
著名算法就是通过递归实现的
注：在大量的递归调用可能会出现栈溢出，但是这个是可以避免并解决这个问题
 */
func main()  {
	/**
		斐波那契数列
	 */
	for i:=0;i<10;i++ {
		fmt.Println(fibo(i))
	}
}
func fibo(n int) int  {
	if n<2 {
		return n
	}
	return fibo(n-2) +fibo(n-1)
}