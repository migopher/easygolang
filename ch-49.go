package main

import (
	"fmt"
	"reflect"
)

/**
结构体
带标签的结构体
 */
type UserTag struct {
	ID       int   `用户ID`
	UserName string `用户名`
	Phone    string `电话`
	int //匿名字段
}

func main() {
	var user UserTag
	user.UserName="goalng"
	user.ID=1
	user.Phone="1555"
	user.int=255
	fmt.Println(user)
	for  i:=0 ; i< 3 ;i++{
		refTag(user,i);
	}
}

/**
反射结构体的tag属性
 */
func refTag(ut UserTag,i int)  {
	utType:=reflect.TypeOf(ut)  //使用反射机制反射
	ixType:=utType.Field(i)
	fmt.Println(ixType.Tag)  	//反射tag内容
}