package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	contents, _ := os.ReadFile("./inputs/10.txt")
	input := string(contents)

	x := 1
	cycles := 0
	signalStrengthSum := 0
	checkCycles := []int{20, 60, 100, 140, 180, 220}
	crt := ""

	cycle := func(limit int) {
		for i := 0; i < limit; i++ {
			position := cycles % 40

			if cycles > 0 && position == 0 {
				crt = crt + "\n"
			}

			if (position == x) || (position == x-1) || (position == x+1) {
				crt = crt + "#"
			} else {
				crt = crt + "."
			}

			cycles++

			for _, value := range checkCycles {
				if value == cycles {
					signalStrengthSum += cycles * x
				}
			}
		}
	}

	for _, row := range strings.Split(input, "\n") {
		parsed := regexp.MustCompile(`(addx|noop)\s?(-?\d+)?`).FindStringSubmatch(row)

		switch parsed[1] {
		case "addx":
			cycle(2)
			value, _ := strconv.Atoi(parsed[2])
			x += value
		case "noop":
			cycle(1)
		}
	}

	fmt.Printf("\nPart 1: %d", signalStrengthSum)
	fmt.Printf("\nPart 2: \n%s", crt)
}
