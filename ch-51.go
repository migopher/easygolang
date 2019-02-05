package main

import "fmt"

type Aa struct {
	a int
	b int
}


func main() {
	var a Aa
	a.a=255
	a.b=5
	fmt.Println(a.Add())
	aa:=Aa{100,200}
	fmt.Println(aa.Add())
}

func (a Aa) Add() int  {
	return a.a +a.b
}

