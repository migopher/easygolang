package main

import "fmt"

/*
指针
 */
func main() {
	a:=24
	fmt.Printf("%x \t",&a)  //输出a变量的内存地址使用&进行访问
	update(&a)
	fmt.Println(a)
	fmt.Println(1)
}

func update(a *int)  {
	var b *int
	b=a  //
	fmt.Println(b)
}