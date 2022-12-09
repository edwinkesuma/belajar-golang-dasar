package main

import "fmt"

func main() {
	nilai32 := 100
	nilai64 := int64(nilai32)
	nilaiFloat64 := float64(nilai64)

	fmt.Printf("Tipe data %T\n", nilai32)
	fmt.Printf("Tipe data %T\n", nilai64)
	fmt.Printf("Tipe data %T\n", nilaiFloat64)

	name := "edwin"
	indexName := name[1]
	convertedIndexName := string(indexName)
	fmt.Println(convertedIndexName)
}
