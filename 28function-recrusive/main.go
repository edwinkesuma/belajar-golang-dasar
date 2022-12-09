package main

import "fmt"

func main() {
	res := forFactorial(5)
	fmt.Println(res)
	fmt.Println(factorialRecursive(5))
}

func forFactorial(value int) int {
	num := 1
	for i := value; i > 0; i-- {
		num *= i
	}
	return num
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}
