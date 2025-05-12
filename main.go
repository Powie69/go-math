package main

import (
	"fmt"
	"github.com/charmbracelet/huh"
	"log"
	"math/rand"
	"strconv"
)

var (
	input   int
	minimum int = 1
	maximum int = 9
)

func main() {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[int]().
				Title("Welcome!").
				Options(
					huh.NewOption("Start", 0),
					huh.NewOption("Set range", 1),
					huh.NewOption("Set seed", 2),
				).
				Value(&input),
		),
	)

	for {
		if err := form.Run(); err != nil {
			log.Fatal(err)
		}
	}
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
