package main

import "fmt"

/**
切片
 */
func main()  {
	/**
	长度为3给元素，容量为5给元素
	 */
	slice:=make([]string,3,5)
	fmt.Println(cap(slice)) //容量
	fmt.Println(len(slice)) //长度
	slice1:=[]string{"java","go","php"}
	fmt.Println(cap(slice1))
	fmt.Println(len(slice1))
	//切片扩容
	newslice1:=append(slice1,"python")
	fmt.Println(newslice1)



}