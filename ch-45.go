package main

import "fmt"

/**
映射
 */
func main() {
	//创建一个映射，键为字符串,值为int型
	where := make(map[string]int)
	where["go"] = 2555
	fmt.Println(where["go"])
	dicc := map[string]int{"golang": 110, "java": 200}
	//查找
	v, exi := dicc["java"]
	//判断键对应的值是否存在
	if exi {
		fmt.Println(v)
	}
	//遍历
	for k, v := range dicc {
		fmt.Printf("Key:%s,Value:%v \n",k,v)
	}
	//指定删除映射里的键删除
	remove(dicc,"java")
	fmt.Println(dicc)
}

func remove(c map[string]int,key string) {
	delete(c,key)
}