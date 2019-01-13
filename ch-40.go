package main

import (
	"fmt"
	"os"
)

/**
defer延迟语句
defer时常用于关闭资源
defer会在return之前执行
 */
func main()  {
	defer func() {  //必须提前执行defer，好捕捉错误
		if info := recover(); info != nil { //recover()用来判断是否有累机错误
			fmt.Println("触发了宕机", info) //如果存在错误打印错误信息
		} else {
			fmt.Println("程序正常退出")     //如果无错运行正常
		}
	}()
	rw:=Rw("1.txt")
	fmt.Println(rw)
}
/**
疯转一个读取文件函数
 */
func Rw(fileName string) string {
	file,err:=os.Open(fileName) //当打开文件失败
	if err !=nil {
		panic(err) //触发一个宕机
	}
	defer file.Close() //
	b1 := make([]byte, 5)
	file.Read(b1)
	return string(b1)
}