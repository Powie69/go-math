package main

import (
	"fmt"
	"github.com/charmbracelet/huh"
	"log"
	"math/rand"
	"os"
	"strconv"
)

var (
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
		guess := askQuestion()

		if guess == "q" {
			break
		}

		fmt.Println(guess)
		fmt.Println(isValidNumber(guess))
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

func askQuestion() (guess string) {
	_, question := makeQuestion()

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title(fmt.Sprintf("What's %s", question)).
				Prompt("= ").
				Value(&guess),
		),
	)

	if err := form.Run(); err != nil {
		log.Fatal(err)
	}

	return guess
}

func makeQuestion() (int, string) {
	first, second := rand.Intn(maximum-minimum)+minimum, rand.Intn(maximum-minimum)+minimum
	return first + second, fmt.Sprintf("%d + %d", first, second)
}

func setRange() {

}

func isValidNumber(s string) (bool, int) {
	result, err := strconv.Atoi(s)
	return err == nil, result
}
