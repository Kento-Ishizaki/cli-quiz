package main

import (
	"testing"
)

func TestNewQuiz(t *testing.T) {
	questions := []Question{
		{
			Question: "What is the capital of France?",
			Choices:  []string{"Paris", "London", "Berlin", "Madrid"},
			Answer:   1,
		},
		{
			Question: "What is the capital of England?",
			Choices:  []string{"Paris", "London", "Berlin", "Madrid"},
			Answer:   2,
		},
	}

	quiz := NewQuiz(questions)

	if len(quiz.Questions) != 2 {
		t.Errorf("Expected 2 questions, but got %d", len(quiz.Questions))
	}

	if quiz.Score != 0 {
		t.Errorf("Expected initial score to be 0, but got %d", quiz.Score)
	}
}

func TestQuiz_Start(t *testing.T) {
	// This test is a bit more complex as it involves user input and output.
	// You might want to use a mocking library to simulate user input and capture output.
	// For now, we'll just test that the function runs without errors with a single question.

	questions := []Question{
		{
			Question: "What is the capital of France?",
			Choices:  []string{"Paris", "London", "Berlin", "Madrid"},
			Answer:   1,
		},
	}

	quiz := NewQuiz(questions)

	// This will not test the user input/output parts of the function.
	// You will need to use a more advanced testing strategy for that.
	quiz.Start()
}
