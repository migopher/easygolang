package main

import "fmt"

/**
使用函数参数进行切换逻辑操作
使用可变参数进行判断最大年龄和最小年龄
 */
type action func(...int) int //

func main() {
	arr := []int{99, 2, 3, 4, 5, 6, 7, 82, 1}
	fmt.Println(getAgeMaxOrMin(min, arr...)) //输出可变参数中最小的值
	fmt.Println(getAgeMaxOrMin(max, arr...)) //输出可变参数中最大值

}
func getAgeMaxOrMin(actFun action, arr ...int) int {
	return actFun(arr...)
}
func min(arrInt ...int) int {
	var min int
	for k, v := range arrInt {
		if min > v || k == 0 { //应min第一次执行前值为0,如可变参数没有0则无法返回数组中的值
			min = v
		}
	}
	return min
}

func max(arrInt ...int) int {
	max := 0
	for _, v := range arrInt {
		if max < v {
			max = v
		}
	}
	return max
}
