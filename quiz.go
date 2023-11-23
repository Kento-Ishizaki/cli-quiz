package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Quiz struct {
	Questions []Question
	Score     int
}

func NewQuiz(questions []Question) *Quiz {
	return &Quiz{
		Questions: questions,
		Score:     0,
	}
}

func (q *Quiz) Start() {
	reader := bufio.NewReader(os.Stdin)

	for i, question := range q.Questions {
		fmt.Printf("Question %d: %s\n", i+1, question.Question)

		for j, choice := range question.Choices {
			fmt.Printf("%d. %s\n", j+1, choice)
		}

		fmt.Print("Enter your answer (1-4): ")
		text, _ := reader.ReadString('\n')
		answer, err := strconv.Atoi(text[:len(text)-1])
		if err != nil {
			fmt.Println("Invalid input. Please enter a number between 1 and 4.")
			i-- // Repeat the current question
			continue
		}

		if answer == question.Answer {
			fmt.Println("Correct!")
			q.Score++
		} else {
			fmt.Println("Incorrect. The correct answer was:", question.Answer)
		}
	}

	fmt.Printf("Quiz finished. Your score: %d/%d\n", q.Score, len(q.Questions))
}
