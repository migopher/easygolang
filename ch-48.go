package main

import "fmt"

/**
结构体与方法

 */

type Users struct {
	ID       int
	UserName string
	Phone    string
}

func main() {
	var user Users
	user.ID = 2
	user.Phone = "110"
	user.UserName = "Go"
	updatePhone(user)
	updateUserName(&user)
	fmt.Println(user)
}

/**
结构体形式传参
传入一个副本对user传入地址没有影响
 */
func updatePhone(user Users) {
	user.Phone = "120"
}

/**
指针传参
影响传递的地址
 */
func updateUserName(user *Users) {
	user.UserName = "golang"
}
