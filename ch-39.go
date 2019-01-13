package main

import "fmt"

/**
内置函数
常用的内置函数
close:用于管道通信
len:用于返回某个类型的长度或数量
cap:返回某个类型的最大容量
new、make:用于内置引用类型
copy、append:用于复制和连接切片
panic、recover：两者均用在错误处理机制上
print、pringtln:用于底层打印函数
简单介绍
 */
func main()  {
	a:="golang"
	fmt.Println(len(a))
	b:=[]int{1,2,3,4}
	println(cap(b))
}