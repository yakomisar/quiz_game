package main

import (
	"encoding/csv"
	"fmt"
	"os"
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

	quiz := transformQuestions(data)
	fmt.Println(quiz)
}
