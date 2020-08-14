package main

import (
	"fmt"
)

func Test() func(i int) int {
	i := 10
	return func(input int) int {
		fmt.Println("i = ", i)
		if input < i {
			i = input
			return -1
		} else {
			i = input
			return 4
		}
	}
}

func main() {
	a := Test()
	fmt.Println(a(4))
	fmt.Println(a(5))
	fmt.Println(a(15))
}
