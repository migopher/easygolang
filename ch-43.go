package main

import "fmt"

/**
多维数组
 */
func main() {
	array := [4][2]int{
		{1, 2},
		{3, 4},
		{4, 5},
		{6, 7},
	}
	fmt.Println(array)
	/**
	修改元素
	 */
	array[0][0]=0
	fmt.Println(array)
	/**
	其他数组赋值
	 */
	var array2  [2]int =array[2]
	fmt.Println(array2)
}
