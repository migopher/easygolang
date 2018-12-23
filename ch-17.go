package main

import "fmt"

/**
数据类型别名
 */
type (
	stringA string
)

func main() {
	var a stringA
	a = "这是中文，"
	b := "这也是中文"
	//fmt.Println(a+b) //因为数据类型别名与原来类型不一样，所以报错了
	fmt.Println(string(a) + b) //强制转换成字符串类型 类型别名就一样了
}
