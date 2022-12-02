package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	Rock     = 1
	Paper    = 2
	Scissors = 3
)

const (
	Loss = 0
	Draw = 3
	Win  = 6
)

func main() {
	contents, _ := os.ReadFile("./inputs/2.txt")
	input := string(contents)

	evaluations := map[string]int{
		"A X": Rock + Draw,
		"A Y": Paper + Win,
		"A Z": Scissors + Loss,
		"B X": Rock + Loss,
		"B Y": Paper + Draw,
		"B Z": Scissors + Win,
		"C X": Rock + Win,
		"C Y": Paper + Loss,
		"C Z": Scissors + Draw,
	}

	// A = rock
	// B = paper
	// C = scissors

	// X = rock
	// Y = paper
	// Z = scissors

	adjustments := map[string]string{
		"A X": "A Z", // lose
		"A Y": "A X", // draw
		"A Z": "A Y", // win
		"B X": "B X", // lose
		"B Y": "B Y", // draw
		"B Z": "B Z", // win
		"C X": "C Y", // lose
		"C Y": "C Z", // draw
		"C Z": "C X", // win
	}

	var scorePart1 int
	var scorePart2 int

	for _, row := range strings.Split(input, "\n") {
		scorePart1 += evaluations[row]
		scorePart2 += evaluations[adjustments[row]]
	}

	fmt.Printf("\nPart 1: %d", scorePart1)
	fmt.Printf("\nPart 2: %d", scorePart2)
}
