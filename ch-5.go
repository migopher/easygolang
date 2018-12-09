package main

import "fmt"

var a=1 //全局变量

func main()  {
	Abc()
	fmt.Println(a)
}
func Abc()  {
	a++
}