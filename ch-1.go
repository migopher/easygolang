package main

import "fmt"

func main() {
	n := 0
	p := &n
	*p++

	fmt.Println(n)
}
