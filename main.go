package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
   // "strconv"
)

func menu() (uint8, error) {

	var choice uint8

	fmt.Println("-- Chose a level of questions --")
	fmt.Println()

	fmt.Println("Level 1")
	fmt.Println("Level 2")
	fmt.Println("Level 3")

	fmt.Scanln(&choice)
	if choice > 3 || choice < 1 {
		return 0, errors.New("Invalid level choice\n")
	}

	return choice, nil

}

type question struct {
	expression string
	answer     int
	difficulty uint8
}

func loadLevel(dif uint8) (bool){

    fmt.Printf("--- Level %d ---\n", dif)

	data, err := os.Open("./db/equations.csv")
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return false
	}
	defer data.Close()

	reader := csv.NewReader(data)

	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return false
	}

    for _, linha := range lines {
        fmt.Println(linha)
    }

    fmt.Println("---------------")

    return true


	return true
}


func main() {

	fmt.Println("----------- welcome to Eureka, a equation quiz! ----------")
	fmt.Println()

	for {

		choice, err := menu()
		if err != nil {
			fmt.Println(err)
			continue
		}

		loadLevel(choice)

		fmt.Println("\nDo you want to play again? (y/n)")
		var playAgain string
		fmt.Scanln(&playAgain)
		if playAgain != "y" {
			break
		}

	}

	fmt.Printf("\nThank you for playing Eureka!\n")
}
