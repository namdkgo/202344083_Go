package main

import "fmt"

func main() {
	a := make([]string, 4, 5)
	a[0] = "a"
	a[2] = "c"
	a[3] = "d"
	as := a[0:2]
	fmt.Println(a, len(a), cap(a))
	as[1] = "z"
	fmt.Println(a, len(a), cap(a))
	fmt.Printf("%x %x %x\n", &a[0], &as[0], &a[1])
}
