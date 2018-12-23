package main

import "fmt"

/*
强制类型转换
 */
func main(){
	intA:=11
	intB:=3
	floatA:=float32(intA)/float32(intB)
	floatB:=intA/intB

	fmt.Println(floatA)
	fmt.Println(floatB)
}