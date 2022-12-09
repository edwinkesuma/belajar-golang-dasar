package main

import "fmt"

func main() {
	ujian := 85
	absensi := 75

	lulusUjian := ujian >= 80
	lulusAbsensi := absensi >= 80

	fmt.Println(lulusUjian)
	fmt.Println(lulusAbsensi)

	lulus := lulusUjian && lulusAbsensi
	fmt.Println(lulus)
}
