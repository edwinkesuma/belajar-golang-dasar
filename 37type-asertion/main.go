package main

import "fmt"

func random() interface{} {
	return "random"
}
func main() {
	res := random()
	resString := res.(string)
	fmt.Println(resString)

}
