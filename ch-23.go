package main

import "fmt"

/**
select语句
比较适合通道(channel)的读写操作
用于多个channel的并发读写
 */
func main()  {
	a:=make(chan int,1024)
	b:=make(chan int,1024)
	/**
	每次输出都不一样
	那个有内容就从那个读
	由于channel的读写操作是阻塞操作
	使用select的可以避免单个的channel的阻塞
	使用default可以避免所有的channel阻塞
	 */
	for i:=0;i<10 ;i++  {
		a<-i
		b<-i
		select {
		case c:=<-a:
			fmt.Println("from a=",c)
		case d:=<-b:
			fmt.Println("from b=",d)
		default:
			fmt.Println("Channel blocking")
		}
	}
}