package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	contents, _ := os.ReadFile("./inputs/6.txt")
	input := string(contents)

	part1 := scanMarker(input, 4)
	part2 := scanMarker(input, 14)

	fmt.Printf("\nPart 1: %d", part1)
	fmt.Printf("\nPart 2: %d", part2)
}

func scanMarker(buffer string, length int) int {
	var position int

	for i := 0; i < len(buffer)-(length-1); i++ {
		isValid := true
		for _, x := range buffer[i : i+length] {
			if strings.Count(buffer[i:i+length], string(x)) > 1 {
				isValid = false
			}
		}
		if isValid {
			position = i + length
			break
		}
	}

	return position
}
