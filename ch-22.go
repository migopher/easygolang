package main

import "fmt"

var x interface{}
/**
switch 初始化语句
 */
func main()  {
	 x=1
	switch a:=x.(type) { //初始化判断类型
	case nil:
		fmt.Println("a=nil",a)
	case int:
		fmt.Println("a=int",a)
	case float64:
		fmt.Println("a=float64",a)

	case string:
		fmt.Println("a=string",a)
	default:
		fmt.Println("未知类型",a)
	}
}