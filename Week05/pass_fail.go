package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	//inputScore, _ := reader.ReadString('\n') 에러 해결법 1
	inputScore, err := reader.ReadString('\n')
	log.Fatal(err) // 에러 해결법 2
	fmt.Println(inputScore)
}
