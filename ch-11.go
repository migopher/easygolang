package main

import (
	"fmt"
	"reflect"
)

func main() {
	var v1 complex64
	v1 = 3.14 + 12i
	v2 := 3.14 + 12i //默认为complex128
	fmt.Println(v1, v2)
	fmt.Println(reflect.TypeOf(v1))
	v3 := complex(3.86, 35) //
	v := v3 + v2
	fmt.Println(v)
	fmt.Println(real(v), imag(v)) //real函数获取实部  imag获取虚部
}
