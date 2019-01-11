package main

import "fmt"

/**
延迟控制语句
defer 用于延迟调用制定函数
只能出现到函数内部
当在函数内部使用多个defer，defer其实是一个栈,遵循原则先入后出

常用于回收资源，清理收尾等工作
 */
func main()  {
	//for i:=0;i<10 ;i++  { 使用defer让for 循环倒序输出
	//	defer fmt.Println(i)
	//}
	defer world()
	defer hell()
	fmt.Print("easy golang ")

}

func hell()  {
	fmt.Print("hello")
}
func world()  {
	fmt.Print(" world!")
}
