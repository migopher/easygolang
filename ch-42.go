package main

import "fmt"

/**
复用数据类型
数组操作
 */
func main() {
	/**
	数组
	 */
	var array1 [5]int //声明5个元素的整形数组
	fmt.Println(array1)
	array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array2)
	array3:=[...]int{1,2,3,4,5,6} //未确定长度可以使用...表示自动判别长度
	fmt.Println(array3)
	array4:=[5]int{1:10,2:20} //制定初始化索引的具体值
	fmt.Println(array4)
	//修改元素
	array4[4]=10
	fmt.Println(array4)

	/**
	数组的值也可以是指针
	 */
	array5:=[5]*int{0:new(int),4:new(int)}
	*array5[0]=1
	array5[2]=new(int)
	*array5[2]=20
	*array5[4]=10
	fmt.Println(array5)
	/**
	字符串数组
	 */
	 string1:=[5]string{"go","lang"}
	 fmt.Println(string1)

	 string2:=[3]*string{
	 	new(string),
	 	new(string),
	 	new(string),
	 }
	*string2[0]="java"
	*string2[1]="go"
	*string2[2]="golang"

	for k,v := range string2{
		fmt.Println(k,*v)
	}
}
