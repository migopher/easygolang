package main

import (
	"fmt"
	"unsafe"
)

/**
工厂方法创建结构体
通过函数实现构造
 */
type File struct {
	fd   int    //文件描述符
	name string //文件名
}

func main() {
	file:=NewFile(10,"./1.txt")
	fmt.Println(unsafe.Sizeof(file)) //unsafe.Sizeof用于知道结构体占用的内存地址
}

func NewFile(fd int,name string) *File {
	if fd<0 {
		return nil
	}
	return &File{fd,name}
}