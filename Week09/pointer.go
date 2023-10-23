package main

import (
	"fmt"
)

func main() {
	a := 10
	pa := &a
	fmt.Printf("%d의 주소는 %p입니다.\n", a, &a)
	fmt.Printf("%p의 값은 %d입니다.\n", pa, *pa)
}
