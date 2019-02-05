package main

import "fmt"

/**
结构体
内嵌结构体，类似于继承
命名冲突，字段不被调用不会包吃，但存在隐患
 */
type A struct {
	Uid  int
	Name string
}

type B struct {
	person A
	ID int
}

func main() {
	b:=B{person:A{255,"username"},ID:255}
	fmt.Println(b)
	//调用person嵌套的Name值
	fmt.Printf(b.person.Name)
	//ID
	fmt.Println(b.ID)
}
