package main

import "fmt"

/**
continue
与break相反用于跳转到指定的代码块
continue仅用于for循环
 */
func main() {
//LOOP1:
//	fmt.Println("continue")
	for i := 0; i < 10; i++ {
		switch  {
		case i>0:
			fmt.Println("LOOP1",i)
			continue
			break //因为switch 有多个满足最后一个满足执行所以使用break跳出switch
		case i==1:
			fmt.Println(i)
		default:
			fmt.Println("其他",i)
		}
		fmt.Println(i)
	}
}
