package main

import "fmt"

func main() {
	var myName [2]string

	myName[0] = "edwin"
	myName[1] = "kesuma"

	fmt.Println(myName)

	values := [4]int{1, 2, 3, 4}
	fmt.Println(values)

	fmt.Println(len(myName))
	fmt.Println(len(values))

}
