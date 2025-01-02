package main

import (
	"fmt"
	"strconv"
	"time"
)

const maxTimePerQuestion = 10

type QuizItem struct {
	Prompt    string
	Choices   [4]string
	Correct   int
}

func poseQuestion(item QuizItem, itemNumber int) int {
	fmt.Printf("\nQuestion %d: %s\n", itemNumber, item.Prompt)
	for idx, choice := range item.Choices {
		fmt.Printf("%d. %s\n", idx+1, choice)
	}

	timer := time.NewTimer(time.Duration(maxTimePerQuestion) * time.Second)
	responseCh := make(chan int)

	go func() {
		var response string
		fmt.Print("Enter your response (1-4) or type 'exit' to quit: ")
		_, err := fmt.Scanln(&response)
		if err != nil {
			responseCh <- -1
			return
		}
		if response == "exit" {
			responseCh <- 999
			return
		}
		parsedResponse, err := strconv.Atoi(response)
		if err != nil || parsedResponse < 1 || parsedResponse > 4 {
			responseCh <- -1
			return
		}
		responseCh <- parsedResponse - 1
	}()

	select {
	case <-timer.C:
		fmt.Println("\nTime's up!")
		return -1
	case response := <-responseCh:
		if response == 999 {
			return 999
		}
		if response < 0 || response > 3 {
			fmt.Println("Invalid response, try again.")
			return -1
		}
		return response
	}
}

func evaluateQuiz(quizItems []QuizItem, userResponses []int) {
	correctCount := 0
	for idx, item := range quizItems {
		if userResponses[idx] == item.Correct {
			correctCount++
		}
	}

	fmt.Printf("\nYou correctly answered %d out of %d questions.\n", correctCount, len(quizItems))
	evaluation := ""
	switch {
	case correctCount == len(quizItems):
		evaluation = "Wow Great Job"
	case correctCount >= len(quizItems)/2:
		evaluation = "Well Done"
	default:
		evaluation = "Needs More Practice"
	}
	fmt.Printf("Performance: %s\n", evaluation)
}

func startQuiz(quizItems []QuizItem) {
	var userResponses []int
	for idx, item := range quizItems {
		response := poseQuestion(item, idx+1)
		if response == -1 {
			fmt.Println("Skipping to the next question due to invalid input.")
			continue
		}

		if response == 999 {
			fmt.Println("Exiting the quiz prematurely...")
			break
		}

		userResponses = append(userResponses, response)
	}

	evaluateQuiz(quizItems, userResponses)
}

func main() {
	quizItems := []QuizItem{
		{
			Prompt:  "What is the capital of Germany?",
			Choices: [4]string{"Berlin", "Madrid", "Paris", "Rome"},
			Correct: 0,
		},
		{
			Prompt:  "what is national sports of America?",
			Choices: [4]string{"Cricket", "Hockey", "Football", "Basketball"},
			Correct: 2,
		},
		{
			Prompt:  "Who is the president of India !!",
			Choices: [4]string{"SRK", "Rahul Gandhi", "Narendra Modi", "Amit Shah"},
			Correct: 1,
		},
	}

	fmt.Println("Welcome to the Interactive Quiz!")
	fmt.Println("Type 'exit' anytime to quit early.")

	startQuiz(quizItems)
}
