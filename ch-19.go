package main

import (
	"fmt"
	"os"
)

var User = os.Getenv("USER") //系统环境变量
func main() {
	a := User == "ixuzl"
	fmt.Println(a)
	fmt.Println(User)
}
