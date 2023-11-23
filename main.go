package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Question struct {
	Question string   `json:"question"`
	Choices  []string `json:"choices"`
	Answer   int      `json:"answer"`
}

func main() {
	// Load questions from JSON file
	questions, err := loadQuestions("questions.json")
	if err != nil {
		fmt.Println("Error loading questions:", err)
		os.Exit(1)
	}

	// Create a new quiz with the loaded questions
	quiz := NewQuiz(questions)

	// Start the quiz
	quiz.Start()
}

func loadQuestions(filename string) ([]Question, error) {
	var questions []Question

	// Read the file
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON data into the questions slice
	err = json.Unmarshal(data, &questions)
	if err != nil {
		return nil, err
	}

	return questions, nil
}
