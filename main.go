package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

func menu() (int, error) {
	var choice int

	fmt.Println("-- Choose a level of questions --")
	fmt.Println()
	fmt.Println("Level 1")
	fmt.Println("Level 2")
	fmt.Println("Level 3")

	_, err := fmt.Scanln(&choice)
	if err != nil {
		return 0, err
	}

	if choice > 3 || choice < 1 {
		return 0, errors.New("Invalid level choice")
	}

	return choice, nil
}

type Questions struct {
	Questions map[string][]Question `json:"questions"`
}

type Question struct {
	ID       int    `json:"id"`
	Equation string `json:"equation"`
	Answer   int    `json:"answer"`
}

func loadLevel(dif int) error {
	difStr := strconv.Itoa(dif)

    fmt.Println()

	fmt.Printf("--- Level %d ---\n", dif)

	file, err := os.Open("questions.json")
	if err != nil {
		return errors.New("Error opening the JSON file")
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return errors.New("Error while reading the JSON file")
	}

	var questions Questions
	if err := json.Unmarshal(bytes, &questions); err != nil {
		return errors.New("Error decoding the JSON")
	}

	qs, ok := questions.Questions[difStr]
	if !ok {
		return fmt.Errorf("No questions found for difficulty %d", dif)
	}

	var score int

	for _, q := range qs {
		var answer int

		fmt.Println()
		fmt.Printf(q.Equation + ": ")

		_, err := fmt.Scanln(&answer)
		if err != nil {
			return errors.New("Invalid input")
		}

		if answer != q.Answer {
			fmt.Printf("Incorrect! The correct answer is %d\n", q.Answer)
		} else {
			fmt.Println("Correct!")
			score++
		}
	}

	fmt.Println("---------------")
	fmt.Printf("You got a score of %d\n", score)

	return nil
}

func main() {
	fmt.Println("----------- Welcome to Eureka, an equation quiz! ----------")
	fmt.Println()

	for {
		choice, err := menu()
		if err != nil {
			fmt.Println(err)
			continue
		}

		err = loadLevel(choice)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("\nDo you want to play again? (y/n)")
		var playAgain string
		fmt.Scanln(&playAgain)
		if playAgain != "y" {
			break
		}
	}

	fmt.Println("\nThank you for playing Eureka!")
}

