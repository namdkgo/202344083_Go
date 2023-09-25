package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int = 9
	var b = 2.7 //실수형 default = float64
	var c, d string = "inha~", "go...."
	var e float32 = 3.14
	var f string
	var g bool
	h := 'Z'
	i := "문자열"
	z, y, x := 1, '2', "3"
	J := "변수명이 대문자로 시작하면 다른 패키지에서 이 변수를 사용할 수 있다."
	koreanZombie := "정찬성" // 카멜 케이스
	// a = 9
	// b = 2.7
	// c = "inha~"
	// d = "go...."
	fmt.Println(float64(a) > b)
	fmt.Println(z, y, x)
	fmt.Println(a, b, c, d, e, f, g, h, i, J, koreanZombie) //f, g = zero value
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(h))
	fmt.Println(reflect.TypeOf(b))

	// fmt.Println(reflect.TypeOf('B'))
	// fmt.Println(reflect.TypeOf(100))
	// fmt.Println(reflect.TypeOf(2.71))
	// fmt.Println(reflect.TypeOf(false))
	// fmt.Println(reflect.TypeOf("Go!"))
	// fmt.Println(math.Floor("inha")) //error(string in float)
	// fmt.Println(strings.Title(3.14)) //error(floor in string)
	// fmt.Println('A', '\n') // rune (unicode)
	// fmt.Println('힣', '가')
	// fmt.Println(math.Floor(2.75))
	// fmt.Println(math.Ceil(2.75))
	// fmt.Println(math.Round(2.75))
	// fmt.Println(strings.Title("오픈소스프로그래밍~\nGo"))
	// fmt.Println(strings.Title("open source programming go!"))
}

// package main

// import "fmt"

// func main() {
// 	println("hello, go!")
// 	fmt.Println("OpenSource Programming~", "Go")
// }
