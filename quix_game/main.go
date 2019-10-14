package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("problems.csv")
	check(err)

	r := csv.NewReader(file)
	correctAnswers := 0
	totalQuestions := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		question := record[0]
		correctAnswer := record[1]
		buf := bufio.NewReader(os.Stdin)
		fmt.Print(fmt.Sprintf("What is %s?", question))
		fmt.Print(">")

		userAnswer, err := buf.ReadString('\n')
		userAnswer = strings.Replace(userAnswer, "\n", "", -1)

		check(err)
		totalQuestions++
		fmt.Printf("user answer is:", userAnswer)
		if userAnswer == correctAnswer {
			correctAnswers++
		}
	}
	fmt.Print(fmt.Sprintf("Congratulations — you got %d questions right out of %d!", correctAnswers, totalQuestions))
}
