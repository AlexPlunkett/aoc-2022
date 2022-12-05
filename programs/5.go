package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	contents, _ := os.ReadFile("./inputs/5.txt")
	input := string(contents)

	parts := strings.Split(input, "\n\n")

	numberOfStacks := len(regexp.MustCompile(`\d`).FindAllString(parts[0], -1))
	stackConfiguration1 := make([][]string, numberOfStacks, numberOfStacks) // part 1
	stackConfiguration2 := make([][]string, numberOfStacks, numberOfStacks) // part 2

	for _, row := range strings.Split(parts[0], "\n") {
		for _, stringIndices := range regexp.MustCompile(`(\[\w\])`).FindAllStringIndex(row, -1) {
			// part 1
			stackConfiguration1[stringIndices[0]/4] = append(stackConfiguration1[stringIndices[0]/4], string(row[stringIndices[0]:stringIndices[1]]))
			// part 2
			stackConfiguration2[stringIndices[0]/4] = append(stackConfiguration2[stringIndices[0]/4], string(row[stringIndices[0]:stringIndices[1]]))
		}
	}

	for _, row := range strings.Split(parts[1], "\n") {
		instructions := regexp.MustCompile(`\d+`).FindAllString(row, -1)

		amount, _ := strconv.Atoi(instructions[0])
		from, _ := strconv.Atoi(instructions[1])
		to, _ := strconv.Atoi(instructions[2])

		// part 1
		for i := 0; i < amount; i++ {
			stackConfiguration1[to-1] = append([]string{stackConfiguration1[from-1][0]}, stackConfiguration1[to-1]...)
			stackConfiguration1[from-1] = append(stackConfiguration1[from-1][:0], stackConfiguration1[from-1][1:]...)
		}

		// part 2
		x := make([]string, amount)
		copy(x, stackConfiguration2[from-1][:amount])
		stackConfiguration2[to-1] = append(x, stackConfiguration2[to-1]...)
		stackConfiguration2[from-1] = append([]string{}, stackConfiguration2[from-1][amount:]...)
	}

	// part 1
	var topCrates1 string
	for _, stack := range stackConfiguration1 {
		topCrates1 = topCrates1 + stack[0]
	}
	topCrates1 = strings.ReplaceAll(topCrates1, "[", "")
	topCrates1 = strings.ReplaceAll(topCrates1, "]", "")

	// part 2
	var topCrates2 string
	for _, stack := range stackConfiguration2 {
		topCrates2 = topCrates2 + stack[0]
	}
	topCrates2 = strings.ReplaceAll(topCrates2, "[", "")
	topCrates2 = strings.ReplaceAll(topCrates2, "]", "")

	fmt.Printf("\nPart 1: %s", topCrates1)
	fmt.Printf("\nPart 2: %s", topCrates2)
}
