package main

import "fmt"

func swap(num1 *int, num2 *int) {
	temp := *num1
	*num1 = *num2
	*num2 = temp
}

func main() {
	a := 10
	b := 20
	var c *int = &a //integer 타입 변수는 integer 타입 포인터 변수에만 담을 수 있음.
	fmt.Printf("%T\n", c)
	d := &b

	fmt.Println(a, b)
	swap(&a, &b)
	fmt.Println(a, b)
	swap(c, d)
	fmt.Println(a, b)
}
