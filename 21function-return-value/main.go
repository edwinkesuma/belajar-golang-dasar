package main

import "fmt"

func main() {
	res := greeting("edwin")
	fmt.Println(res)
}

func greeting(name string) string {
	if name != "" {
		return "Hi " + name
	} else {
		return "Hi what's ur name?"
	}
}
