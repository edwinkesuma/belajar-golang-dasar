package main

import "fmt"

func main() {
	name := "edwin"
	greet := func() {
		name := "kesuma"
		fmt.Println(name)
	}
	greet()
	fmt.Println(name)
}
