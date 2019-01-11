package main

import "fmt"

/**
goto语句
goto语句可以无条件跳转到相同函数中的带标签语句
 */
func main()  {
//GG:  //GG标签如果放到for 前面会一直死循环，因为goto一直在跳转
//	fmt.Println("goto stop")
//	hello()
	for i:=0;i<10;i++ {
		fmt.Println(i)
		if i==5 {
			goto GG
		}
	}
GG:  //GG标签如果放到for 前面会一直死循环，因为goto一直在跳转
	fmt.Println("goto stop")
	hello()

}
func hello()  {
	fmt.Println("hello world")
}