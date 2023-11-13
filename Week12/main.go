package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := []int{0, 0, 0, 0, 0} //slice 리터럴
	s[4] = 99
	s[2] = 11
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
	fmt.Println("  __  ")
	copyS := s[1:5]
	for i := 0; i < len(copyS); i++ {
		fmt.Println(copyS[i])
	}

	fmt.Println(reflect.TypeOf(copyS))

	test := [3]string{"inha", "go", "student"}
	testS := test[:2]
	testS2 := test[1:]
	test[1] = "1"
	fmt.Println(testS[1])
	fmt.Println(testS2[0])
}
