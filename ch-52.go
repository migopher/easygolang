package main

import "fmt"

/**
非结构体
 */
type Arr []int

func main() {
	a := Arr{1, 2, 3, 4}.Sum()
	fmt.Println(a)
}

func (v Arr) Sum() (s int) {
	var varlue int
	for _, x := range v {
		varlue += x
	}
	return varlue
}
