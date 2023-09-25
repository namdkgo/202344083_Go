package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("Enter a score: ")
	reader := bufio.NewReader(os.Stdin)
	inputScore, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	if inputScore >= 90 { //invalid operation: inputScore >= 90 (mismatched types string and untyped int)
		grade := "A"
	} else {
		grade := "under A grade"
	}

	fmt.Println(grade)
	fmt.Println(inputScore)
}
