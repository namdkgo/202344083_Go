package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a score: ")
	reader := bufio.NewReader(os.Stdin)
	inputScore, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	inputScore = strings.TrimSpace(inputScore)
	score, err := strconv.ParseFloat(inputScore, 64)
	if err != nil {
		log.Fatal(err)
	}

	var grade string

	if score >= 90 {
		grade = "A"
	} else {
		grade = "under A grade"
	}

	fmt.Println(inputScore)
	fmt.Println("you will get", grade)
}
