package main

import "fmt"

func main() {
	s := "hello "
	fmt.Println(s)
	fmt.Println(len(s)) //字符串长度
	for _, v := range s { //遍历循环字符串
		fmt.Println(string(v))
	}
	s += "world!"
	for i := 0; i < len(s); i++ {
		fmt.Println(string(s[i])) //使用string，否则为字节值
	}
	fmt.Println(s)
	fmt.Println("---修改字符串中都字节---")
	b := []byte(s)
	fmt.Println(b)
	b[5] = ','
	fmt.Printf("%s\n", b)
	fmt.Println("---修改字符串中字符---")
	c := []rune(s)
	c[6] = ','
	c[11] = '。'
	fmt.Println(string(c))
}
