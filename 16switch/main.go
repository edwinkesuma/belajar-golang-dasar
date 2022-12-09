package main

import "fmt"

func main() {
	name := "edwin"

	switch name {
	case "edwin":
		fmt.Println("Hello edwin!")
	case "kesuma":
		fmt.Println("Hello kesuma!")
	default:
		fmt.Println("siapa kamu?")
	}

	//switch dengan short condition
	switch len(name) < 5 {
	case true:
		fmt.Println("nama kurang dari 5 karakter")
	case false:
		fmt.Println("nama sama dengan atau lebih dari 5 karakter")
	}

	length := len(name)
	//switch tanpa kondisi
	switch {
	case length > 5:
		fmt.Println("nama lebih dari 5 karakter")
	case length < 5:
		fmt.Println("nama kurang dari atau sama dengan 5 karakter")
	}
}
