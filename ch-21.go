package main

import (
	"fmt"
	"strconv"
)

/**
选择语句 switch case
 */
func main() {
	a := 30
	switch a {
	case 18:
		fmt.Println("my age:" +strconv.Itoa(a))
	case 24:
		fmt.Println("my age:" +strconv.Itoa(a))
	default:
		fmt.Println("my age:" +strconv.Itoa(a))
	}
	switch  {
	case a<18:
		fmt.Println("少年")
	case a>=18 && a<30:
		fmt.Println("青年")
	case a>=30 && a<40:
		fmt.Println("中年")
	default:
		fmt.Println("老年")
	}
}
