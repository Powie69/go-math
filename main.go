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

type MenuAction int

var (
	isRangeToggled  bool
	ranges          []int
	menuPromptInput MenuAction
	minimum         = 1
	maximum         = 9
)

const (
	StartGame MenuAction = iota
	SetRange
	ToggleRange
	SetBounds
	Quit
)

func main() {
	for {
		makeMenuPrompt()

		switch menuPromptInput {
		case StartGame:
			gameLoop()
		case SetRange:
			setRange()
		case ToggleRange:
			isRangeToggled = !isRangeToggled
		case SetBounds:
			setBounds()
		case Quit:
			fmt.Println("Goodbye!")
			os.Exit(0)
		}
	}
}

func gameLoop() {
	if isRangeToggled && len(ranges) == 0 {
		setRange()
	}

	for {
		answer, guess, questionString := askQuestion()

		if guess == "" {
			fmt.Println("Enter something")
			continue
		}

		if guess == "q" || guess == "x" || guess == "l" || guess == "c" {
			break
		}

		isValid, guessInt := isValidNumber(guess)
		if !isValid {
			fmt.Printf("not a valid number %s\n", guess)
		}
		if guessInt == answer {
			fmt.Println("correct")
			continue
		}
		fmt.Printf("> %s = %d\n", questionString, answer)
	}
}

func makeMenuPrompt() {
	if err := huh.NewSelect[MenuAction]().
		Title("Welcome!").
		Options(
			huh.NewOption("Start", StartGame),
			huh.NewOption("Set range", SetRange),
			huh.NewOption("Toggle range", ToggleRange),
			huh.NewOption("Set Bounds", SetBounds),
			huh.NewOption("Quit", Quit),
		).
		Description(fmt.Sprintf("Range: %s", map[bool]string{true: "✔️", false: "✖️"}[isRangeToggled])).
		Value(&menuPromptInput).
		Run(); err != nil {
		log.Fatal(err)
	}
}

func askQuestion() (answer int, guess string, questionString string) {
	answer, question := makeQuestion()

	if err := huh.NewInput().
		Title(fmt.Sprintf("\nWhat's %s", question)).
		Prompt("= ").
		//Validate(validateNumber).
		Value(&guess).
		Run(); err != nil {
		log.Fatal(err)
	}

	return answer, guess, question
}

func makeQuestion() (answer int, question string) {
	if !isRangeToggled {
		first, second := rand.Intn(maximum-minimum+1)+minimum, rand.Intn(maximum-minimum)+minimum
		return first + second, fmt.Sprintf("%d + %d", first, second)
	}
	first, second := ranges[rand.Intn(len(ranges))], rand.Intn(maximum-minimum+1)+minimum
	if rand.Intn(2) == 0 {
		return first + second, fmt.Sprintf("%d + %d", first, second)
	} else {
		return first + second, fmt.Sprintf("%d + %d", second, first)
	}
}

func setRange() {
	userInput := ""

	if err := huh.NewInput().
		Title("Select Ranges").
		Placeholder(strings.Trim(fmt.Sprint(ranges), "[]")).
		//Validate(validateNumber).
		Value(&userInput).
		Run(); err != nil {
		log.Fatal(err)
	}

	if userInput == "" {
		fmt.Println("enter something idiot")
		return
	}

	args := strings.Fields(userInput)
	intRanges := make([]int, len(args))

	for i, s := range args {
		valid, value := isValidNumber(s)
		if !valid {
			fmt.Printf("%q is not a valid value\n", s)
			return
		}
		intRanges[i] = value
	}

	if !isRangeToggled {
		isRangeToggled = true
	}

	ranges = intRanges
	fmt.Printf("Ranges set to: %v\n", ranges)
}

func setBounds() {
	var minInput, maxInput string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Set minimum").
				Placeholder(strconv.Itoa(minimum)).
				Validate(validateNumber).
				Value(&minInput),
			huh.NewInput().
				Title("Set maximum").
				Placeholder(strconv.Itoa(maximum)).
				Validate(validateNumber).
				Value(&maxInput),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	newMin, newMax := minimum, maximum

	if val, err := strconv.Atoi(minInput); err == nil {
		newMin = val
	}
	if val, err := strconv.Atoi(maxInput); err == nil {
		newMax = val
	}

	if newMin >= newMax {
		fmt.Printf("minimum must be less than maximum\n")
		return
	}

	minimum, maximum = newMin, newMax
}

func validateNumber(s string) error {
	if s == "" {
		return nil
	}
	if _, err := strconv.Atoi(s); err != nil {
		return fmt.Errorf("must be a number")
	}
	return nil
}

func isValidNumber(s string) (bool, int) {
	result, err := strconv.Atoi(s)
	return err == nil, result
}
