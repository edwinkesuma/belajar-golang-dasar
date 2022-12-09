package main

import "fmt"

func main() {
	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	mySlice := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("days capasity", cap(days))
	fmt.Println("myslice capasity", cap(mySlice))
	fmt.Println(days)
	fmt.Println(mySlice)

	slice1 := days[5:]
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	fmt.Println(slice1)

	//mengubah isi slice yang berasal dari array akan mengubah isi array. begitu juga sebaliknya
	slice1[0] = "boo"
	fmt.Println(slice1)
	fmt.Println(days)

	slice2 := append(days[5:])
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	slice3 := append(slice2, "baa")
	slice3[1] = "hehe"
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice4 := append(days[:3])
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

	slice4[1] = "haha"
	fmt.Println(slice4)
	fmt.Println(days)

}
