package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

type questions struct {
	question string
	answer   string
}

func transformQuestions(data [][]string) []questions {
	var questionsList []questions
	for _, val := range data {
		var question questions
		question.question = val[0]
		question.answer = val[1]
		questionsList = append(questionsList, question)
	}
	return questionsList
}

func stopGame(endGame chan<- bool) {

}

func main() {
	gameOver := make(chan bool)
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

	quiz := transformQuestions(data)
	for index, q := range quiz {
		fmt.Printf("Problem #%d: %v\n", index, q.question)

		<-time.After(time.Second * 3)
	}

	select {
	case <-gameOver:
		fmt.Printf("You scored %d out of 999\n", result)
	case <-time.After(4 * time.Second):
		fmt.Println("There's no more time to this. Exiting!")
	}
}
