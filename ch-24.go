package main

import "fmt"

/**
for循环
在go语言中没有while关键字
 */
func main()  {
	/**
	for 循环语句
	 */
	for i:=0;i<10 ;i++  {
		fmt.Println(i)
	}

	a:=0
	b:=5
	/**
	for 条件子语句
	当a为false停止循环
	 */
	for a<b  {
		a++
		fmt.Println(a)
	}
	/**
	for range子语句
	作用类似迭代器
	用于轮询数组或切片值
	字符串值v返回为run类型可以使用string()强制转换字符串
	可以支持字符串,数组，切片，字段，甚至通道等
	比较常用
	*/
	 str:="easygolang"
	for k,v:=  range str{
		fmt.Println("No:",k,"value:",string(v))
	}


}