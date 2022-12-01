package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	contents, _ := os.ReadFile("./inputs/1.txt")
	input := string(contents)

	var sumOfCalories []int
	var runningCalories int

	for _, calories := range strings.Split(input, "\n") {
		if calories == "" {
			sumOfCalories = append(sumOfCalories, runningCalories)
			runningCalories = 0
		} else {
			toInteger, _ := strconv.Atoi(calories)
			runningCalories += toInteger
		}
	}

	sort.Ints(sumOfCalories)

	var sumOfHighestThree int
	for _, calories := range sumOfCalories[len(sumOfCalories)-3:] {
		sumOfHighestThree += calories
	}

	fmt.Printf("\nPart 1: %d", sumOfCalories[len(sumOfCalories)-1])
	fmt.Printf("\nPart 2: %d", sumOfHighestThree)
}
