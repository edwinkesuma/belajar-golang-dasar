package main

import "fmt"

func main() {
	//type declaration digunakan untuk membuat alias data type
	type married bool
	var isMarried married = false

	fmt.Printf("Is married? %v\n", isMarried)
}
