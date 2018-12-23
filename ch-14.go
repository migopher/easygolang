package main

import (
	"fmt"
	"reflect"
	"strconv"
)

/**
strconv包
字符串与其他类型的转换
所有类型几乎可以转换成字符串
其他类型转换字符串不一定转换成功
 */
func main() {
	//var  strA string ="h"
	var intB string = "23 t3"
	fmt.Println(strconv.IntSize) //可用于判断操作系统的位数
	num, _ := strconv.Atoi(intB) //将字符串转换成整数
	fmt.Println(reflect.TypeOf(num))
	str := strconv.Itoa(233); //将整数转换成字符串
	fmt.Println(reflect.TypeOf(str))
	strFloat:=strconv.FormatFloat(3.1415926,'f',-1,64) //浮点数转换成字符串
	fmt.Println(reflect.TypeOf(strFloat),strFloat)
	floatA,_:=strconv.ParseFloat("3.1415926",64) //字符串转换成浮点型64
	fmt.Println(reflect.TypeOf(floatA),floatA)
	fmt.Println(strconv.IsPrint(rune('@')))   //一个字符是否是可打印的
	fmt.Println(strconv.CanBackquote(intB))   //判断字符串是否有反符号`
	fmt.Println(strconv.Quote(`"Fran & Freddie‘s '/' Diner ?"`)) //对双引号字面值表示，控制字符、不可打印字符会进行转义
	s := strconv.QuoteRune(05) //控制字符、不可打印字符会进行转义
	fmt.Println(s)

}
