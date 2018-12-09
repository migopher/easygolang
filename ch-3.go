package main

import "fmt"

func main()  {
	a:=20
	b:=200
	b=a
	fmt.Println(b)
	b+=a
	fmt.Println(b)
	b-=a
	fmt.Println(b)
	b*=a
	fmt.Println(b)
	b/=a
	fmt.Println(b)
	b<<=2
	fmt.Println(b)
	b>>=2
	fmt.Println(b)
	b&=2
	fmt.Println(b)
	b^=2
	fmt.Println(b)
	b|=2
	fmt.Println(b)
}