package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello, this is main")

	stack := make([]int, 0)
	stack = append(stack, 111, 222, 333, 444)
	v := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	len1 := len(stack)
	fmt.Println("stack len = ", len1)
	fmt.Println(v)

	if len(stack) != 0 {
		fmt.Println("123")
		fmt.Println(len(stack))
	}

}
