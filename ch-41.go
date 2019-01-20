package main

import (
	"fmt"
	"os"
)

/**
错误与恢复
go没有像java提供一场机制，所以go提供了panic和recover机制处理异常
panic相当java中的异常抛出，使程序中止，合理使用recover处理panic
合理使用panic
 */

type PathError struct {
	Op string
	Path string
	Err error
}
func main()  {
	_,err:=Stat("12.txt")
	if err !=nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println(200)
	}
	//defer selerror() //研究错误信息
	_,err=os.Open("1123.txt")
	defer selerror() //排除恐慌状态
	if err != nil {
		panic(500) //发出恐慌，二次触发恐慌
	}
}

func (e *PathError) Error() string  {
	return e.Op+" "+e.Path+" "+e.Err.Error()
}
func Stat(name string)(fi *os.File,err error)  {
	defer selerror() //排除恐慌状态,程序继续执行
	file,erro:=os.Open(name)
	if erro !=nil {
		panic(404)
	}
	return file,nil
}
func selerror()  {
	err:=recover()
	if err!=nil {
		fmt.Println(err) //拦截显示错误
		return
	}
}