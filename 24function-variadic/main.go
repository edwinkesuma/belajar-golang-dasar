package main

import "fmt"

func main() {
	numb := []int{10, 10, 10, 10}
	variadicFunction("hi", 10, 12, 21, 32, 54)
	variadicFunction("hello", numb...)
}

func variadicFunction(greet string, number ...int) {
	x := 0
	for _, i := range number {
		x += i
	}
	fmt.Println(greet, x)
}
