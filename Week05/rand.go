package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	answerNum := rand.Intn(100) + 1
	reader := bufio.NewReader(os.Stdin)

	for i := 10; i > 0; i-- {
		fmt.Println("남은 기회 :", i)
		fmt.Print("입력 : ")

		inputNumber, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		inputNumber = strings.TrimSpace(inputNumber)
		inputNum, err := strconv.Atoi(inputNumber)
		if err != nil {
			log.Fatal(err)
		}

		if inputNum > answerNum {
			fmt.Println("입력한 값보다 작습니다.")
		} else if inputNum < answerNum {
			fmt.Println("입력한 값보다 큽니다.")
		} else {
			fmt.Println(inputNum, "정답입니다!")
			break
		}

		if i == 1 {
			fmt.Println("실패!")
		}
	}

	fmt.Println(answerNum)
}
