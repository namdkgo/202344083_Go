package main

import (
	"fmt"
	"strings"
)

func main() {
	brokenWords := "cs r?cks~"
	replacer := strings.NewReplacer("?", "o")
	fixedWords := replacer.Replace(brokenWords)
	fmt.Println(fixedWords)
}
