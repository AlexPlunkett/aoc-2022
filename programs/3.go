package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

const Priorities = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	contents, _ := os.ReadFile("./inputs/3.txt")
	input := string(contents)

	var sumOfPrioritiesPart1 int
	var sumOfPrioritiesPart2 int

	rows := strings.Split(input, "\n")

	for n, row := range rows {
		// part 1

		a := row[:len(row)/2] // left
		b := row[len(row)/2:] // right

		for i := 0; i < len(row)/2; i++ {
			if strings.Contains(b, string(a[i])) {
				sumOfPrioritiesPart1 += strings.Index(Priorities, string(a[i])) + 1
				break
			}
		}

		// part 2

		if n%3 == 0 {
			group := []string{rows[n], rows[n+1], rows[n+2]}

			sort.SliceStable(group, func(i, j int) bool {
				return len(group[j]) < len(group[i])
			})

			for _, x := range group[0] {
				if strings.Contains(group[1], string(x)) && strings.Contains(group[2], string(x)) {
					sumOfPrioritiesPart2 += strings.Index(Priorities, string(x)) + 1
					break
				}
			}
		}
	}

	fmt.Printf("\nPart 1: %d", sumOfPrioritiesPart1)
	fmt.Printf("\nPart 2: %d", sumOfPrioritiesPart2)
}
