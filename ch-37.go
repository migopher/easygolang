package main

import "fmt"

/**
闭包函数案例
 */
func main()  {
	a:=5
	f:= func() func() int {
		return func() int {
			return a+4
		}
	}()
	fmt.Println(f()) //9
	a=10
	fmt.Println(f()) //14
}