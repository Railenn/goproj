package main

import (
	"fmt"
	"math/rand"
)
func difficultySelector()  {
	fmt.Print(
		"Please select the difficulty level: \n",
		"1. Easy (10 chances) \n",
		"2. Medium (5 chances) \n",
		"3. Hard (3 chances) \n\n",
		"Enter your choice: ",
	)
}

func defineChances(difficultyLevel int) int {
	if (difficultyLevel == 1) {
		return 10
	}
	if (difficultyLevel == 2) {
		return 5
	}
	if (difficultyLevel == 3) {
		return 3
	}
	return 0
}

func defineDifficultyLevelName(difficultyLevel int) (difficultyLevelName string) {
	switch difficultyLevel {
	case 1:
		difficultyLevelName = "Easy"
	case 2:
		difficultyLevelName = "Medium"
	case 3:
		difficultyLevelName = "Hard"
	}
	return difficultyLevelName
}

func repeatGuessing() bool {
	var answer int
	fmt.Print(
		"\nDo you want to try guessing again?\n",
		
		"1. Yes\n",
		"2. No\n",
		"Enter your choice: ",
	)
	fmt.Scanln(&answer)
	fmt.Println()
	if answer == 1 {
		return true
	}
	return false
}
	
func main() {
	const firstNumber int = 1
	const lastNumber int = 100
    
	var difficultyLevel int = 0

	var secretNumber int
	var playersGuess int
	

	fmt.Println(
		"Welcome to the Number Guessing Game!\n",
		"I'm thinking about number between", 
		firstNumber, "and", lastNumber, "\n",
	)

	for {	
		difficultyLevel = 0

		for difficultyLevel < 1 || difficultyLevel > 3 {
			difficultySelector()
			fmt.Scanln(&difficultyLevel)
		}
		
		secretNumber = rand.Intn(lastNumber+1)
		

		fmt.Println(
			"\nGreat! You have selected the", defineDifficultyLevelName(difficultyLevel), 
			"difficulty level.",
			"Let's start the game!\n",
		)
		
		for i := 0; i < defineChances(difficultyLevel); i++ {
			fmt.Print("Enter your guess: ")
			fmt.Scanln(&playersGuess)
			
			if playersGuess > secretNumber {
				fmt.Println("Incorrect! The number is less than", playersGuess)
			}
			if playersGuess < secretNumber {
				fmt.Println("Incorrect! The number is greater than", playersGuess )
			}
			if playersGuess == secretNumber {
				fmt.Println(
					"Congratulations! You guessed the correct",
					"number in", (i+1), "attempts.",
				)
				break
			}
		}

		if repeatGuessing() {
			continue
		} else {
			break
		}
	}
}

