package greeting

import "fmt"

func Hello() { //소문자로 시작하면 여기서만 사용 가능, 대문자로 만들면 공유됨
	fmt.Println("안녕하세요!")
}

func Hi() {
	fmt.Println("안녕!")
}
