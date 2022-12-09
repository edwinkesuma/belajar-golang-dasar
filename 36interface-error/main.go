package main

import (
	"errors"
	"fmt"
)

func main() {
	res, err := pembagi(5, 0)

	if err == nil {
		fmt.Printf("Hasil: %v\n", res)
	} else {
		fmt.Println(err)
	}
}

func pembagi(angka int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh null")
	} else {
		return angka / pembagi, nil
	}
}
