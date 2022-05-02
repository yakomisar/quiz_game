package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

type questions struct {
	question    string
	answer      string
	user_answer string
	is_correct  bool
}

func transformQuestions(data [][]string) []questions {
	var questionsList []questions
	for _, val := range data {
		var question questions
		question.question = val[0]
		question.answer = val[1]
		question.user_answer = ""
		question.is_correct = false
		questionsList = append(questionsList, question)
	}
	return questionsList
}

func main() {
	file, err := os.Open("problems.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}
	stopGame := time.NewTimer(time.Second * 10)
	quiz := transformQuestions(data)
	correct := 0
	for index, q := range quiz {
		fmt.Printf("Problem #%d: %v = ", index, q.question)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scan(&answer)
			answerCh <- answer
		}()
		select {
		case <-stopGame.C:
			fmt.Printf("\nTime is over. You scored %d out of %d", correct, len(quiz))
			return
		case q.user_answer = <-answerCh:
			if q.user_answer == q.answer {
				correct++
			}
		}
	}
	fmt.Printf("You scored %d out of %d", correct, len(quiz))
}
