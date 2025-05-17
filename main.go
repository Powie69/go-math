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
	isRangeToggled bool
	ranges         []int
	input          int
	minimum        int = 1
	maximum        int = 9
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
			isRangeToggled = !isRangeToggled
		case 3:
			fmt.Println("Goodbye!")
			os.Exit(0)
		}
	}
}

func gameLoop() {
	for {
		answer, guess, questionString := askQuestion()

		if guess == "" {
			println("Enter something idiot")
		}

		if guess == "q" || guess == "x" || guess == "l" {
			break
		}

		isValid, guessInt := isValidNumber(guess)
		if !isValid {
			fmt.Printf("not a valid number %s\n", guess)
			continue
		}
		if guessInt == answer {
			println("correct")
			continue
		}
		fmt.Printf("WRONG!\n d %s = %d\n", questionString, answer)
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
					huh.NewOption("Toggle range", 2),
					huh.NewOption("Quit", 3),
				).
				Description(fmt.Sprintf("Range: %s", map[bool]string{true: "✔️", false: "✖️"}[isRangeToggled])).
				Value(&input),
		),
	)

	if err := form.Run(); err != nil {
		log.Fatal(err)
	}
}

func askQuestion() (answer int, guess string, questionString string) {
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

	return answer, guess, question
}

func makeQuestion() (answer int, question string) {
	if isRangeToggled {
		first, second := ranges[rand.Intn(len(ranges))], rand.Intn(maximum-minimum)+minimum
		return first + second, fmt.Sprintf("%d + %d", first, second)
	}
	first, second := rand.Intn(maximum-minimum)+minimum, rand.Intn(maximum-minimum)+minimum
	return first + second, fmt.Sprintf("%d + %d", first, second)
}

// uncomment when 4+ uses
//func toggleRange() {
//	isRangeToggled = !isRangeToggled
//}

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

	if !isRangeToggled {
		isRangeToggled = true
	}

	ranges = intRanges
	fmt.Printf("Ranges set to: %v\n", ranges)
}

func isValidNumber(s string) (bool, int) {
	result, err := strconv.Atoi(s)
	return err == nil, result
}
