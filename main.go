package main

import (
	"fmt"
	"github.com/charmbracelet/huh"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var (
	ranges  = []int{}
	input   int
	minimum int = 1
	maximum int = 9
)

func main() {
	for {
		makePrompt()

		switch input {
		case 0:
			gameLoop()
		case 1:
			setRange()
		case 2:
			fmt.Println("Goodbye!")
			os.Exit(0)
		}
	}
}

func gameLoop() {
	for {
		answer, guess := askQuestion()

		if guess == "q" {
			break
		}

		fmt.Println(guess)
		fmt.Println(isValidNumber(guess))
		isValid, guessInt := isValidNumber(guess)
		guess = ""
		if !isValid {
			fmt.Printf("not a valid number  %s", guess)
			continue
		}
		if guessInt == answer {
			println("correct")
			continue
		}
		println("WRONG!")
	}
}

func makePrompt() {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[int]().
				Title("Welcome!").
				Options(
					huh.NewOption("Start", 0),
					huh.NewOption("Set range", 1),
					huh.NewOption("Quit", 2),
				).
				Value(&input),
		),
	)

	if err := form.Run(); err != nil {
		log.Fatal(err)
	}
}

func askQuestion() (answer int, guess string) {
	answer, question := makeQuestion()

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title(fmt.Sprintf("What's %s", question)).
				Prompt("= ").
				Value(&guess),
		),
	)

	err := form.Run()

	if err != nil {
		log.Fatal(err)
	}

	return answer, guess
}

func makeQuestion() (int, string) {
	first, second := rand.Intn(maximum-minimum)+minimum, rand.Intn(maximum-minimum)+minimum
	return first + second, fmt.Sprintf("%d + %d", first, second)
}

func setRange() {
	userInput := ""

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Select Ranges").
				Value(&userInput),
		),
	)

	err := form.Run()

	if err != nil {
		log.Fatal(err)
	}

	if userInput == "" {
		fmt.Println("enter something idiot")
		return
	}

	args := strings.Fields(userInput)
	intRanges := make([]int, len(args))

	for i, s := range args {
		valid, val := isValidNumber(s)
		if !valid {
			fmt.Printf("%q is not a valid value\n", s)
			return
		}
		intRanges[i] = val
	}

	ranges = intRanges
	fmt.Printf("Ranges set to: %v\n", ranges)
}

func isValidNumber(s string) (bool, int) {
	result, err := strconv.Atoi(s)
	return err == nil, result
}
