package main

import (
	"Week10/src/greeting"
	"Week10/src/mymath"
	"fmt"
)

func main() {
	greeting.Hello()
	greeting.Hi()
	fmt.Println(mymath.Myabs(-5))
	fmt.Println(mymath.Myabs(5))
	fmt.Println(mymath.Mypow(2, 10))
	fmt.Println(mymath.Mypow(3, 3))
	fmt.Println(mymath.Mypow(10, 0))
	fmt.Println(mymath.Mypow(2, -3))
}
