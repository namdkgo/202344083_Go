package main

import "fmt"

func main() {
	var a []string
	var b []bool
	//a = make([]string, 4, 5)
	a = append(a, "abc")
	b = append(b, true)
	fmt.Printf("%#v %#v\n", a, b)
	fmt.Println(a, len(a), cap(a))
	fmt.Println(b, len(b), cap(b))
}
