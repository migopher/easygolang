package main

import "fmt"

/**
流程控制
if判断
if-else
else-if
 */
func main() {
	b := 1
	if b == 1 {
		fmt.Println("相等")
	}
	a := 21
	if a < 20 {
		fmt.Println("a小于20")
	} else {
		fmt.Println("a大于二十")
	}
	c := 0
	if c > 20 {
		fmt.Println("c大于20")
	} else if c < 20 {
		fmt.Println("c小于20")
	} else {
		fmt.Println("c等于20")
	}

	//初始化子语句
	if x:=3.14 ;x == 3.14 {
		fmt.Println(x)
	}

}
