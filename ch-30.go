package main

import "fmt"

/**
函数
main()函数它是程序的入口文件
使用func进行声明函数
函数名如果是小写开头作用域只属于所声明的包
如果是大写则该函数公开，可以被其他包调用
大小写适合所有变量、函数

函数名小写开头为私有方法
函数名大写开头为共有方法类似
go语言有比较先进的特效 比如:多返回值,延迟语句(defer),可变参数等
*/
/**
声明一个函数类型
 */
type addfun func(int, int) int

func main() {
	var_add := add             //将add函数赋值给var_add
	fmt.Println(var_add(1, 2)) //直接使用变量函数类型调用add函数
	fmt.Println(take(2019, 1))
	fmt.Println(Max(2019,2018));
}

/**
定义一个add函数传a和b参数类型为整形，进行计算返回一个整形的计算结果
返回 返回可以返回多值
 */
func add(a, b int) int { //小写开头为私有方法
	return a + b
}

/**
定义take函数
传入a和b参数进行相乘计算
返回3个参数
返回 a参数与b参数与a和b想成的参数进行返回
 */
func take(a, b int) (int, int, int) {
	return a, b, a * b
}

func Max(x,y int) int {
	if x>y{
		return x
	}
	return y
}
