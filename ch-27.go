package main

import "fmt"

/**
break语句
用于打断当前流程控制
break可以用于跳出无限循环
 */
func main()  {
	for i:=0;i<10;i++ {
		for j:=0;j<10;j++  {
			if(i == j){
				break  //跳出里面j for循环
			}
			fmt.Println(j)
		}
	}
	/**
	用于跳出for 无限循环，避免造成死循环
	 */
	x:=0
	for{
		fmt.Println("x=",x)
		if x == 5 {
			break
		}
		x++
	}
	
}