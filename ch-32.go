package main

import "fmt"

/**
将函数作为参数
对函数进行各种逻辑切换，使程序更灵活
 */
type isBool func(int) bool

func main() {
	arrInt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	a := filter(arrInt, isOdd) //过滤所有的偶数
	fmt.Println(a)
	b := filter(arrInt, isEven) //过滤所有的奇数
	fmt.Println(b)
}

func filter(arr []int, b isBool) []int {
	var resu []int
	for _, v := range arr {
		if b(v) {
			resu = append(resu, v);
		}
	}
	return resu
}

func isOdd(v int) bool {
	if v%2 == 0 {
		return false
	}
	return true
}

func isEven(v int) bool {
	if v%2 == 0 {
		return true
	}
	return false
}
