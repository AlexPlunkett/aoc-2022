package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	contents, _ := os.ReadFile("./inputs/4.txt")
	input := string(contents)

	var totalContainments int
	var totalOverlaps int

	for _, row := range strings.Split(input, "\n") {
		pair := strings.Split(row, ",")

		// part 1

		sort.SliceStable(pair, func(i, j int) bool {
			left := strings.Split(pair[j], "-")
			right := strings.Split(pair[i], "-")

			a, _ := strconv.Atoi(left[0])
			b, _ := strconv.Atoi(left[1])
			c, _ := strconv.Atoi(right[0])
			d, _ := strconv.Atoi(right[1])

			return (b - a) < (d - c)
		})

		left := strings.Split(pair[0], "-")
		right := strings.Split(pair[1], "-")

		a, _ := strconv.Atoi(left[0])
		b, _ := strconv.Atoi(left[1])
		c, _ := strconv.Atoi(right[0])
		d, _ := strconv.Atoi(right[1])

		if a <= c && b >= d {
			totalContainments++
		}

		// part 2

		sort.SliceStable(pair, func(i, j int) bool {
			left := strings.Split(pair[j], "-")
			right := strings.Split(pair[i], "-")

			a, _ := strconv.Atoi(left[0])
			c, _ := strconv.Atoi(right[0])

			return a > c
		})

		left = strings.Split(pair[0], "-")
		right = strings.Split(pair[1], "-")

		a, _ = strconv.Atoi(left[0])
		b, _ = strconv.Atoi(left[1])
		c, _ = strconv.Atoi(right[0])
		d, _ = strconv.Atoi(right[1])

		if b >= c {
			totalOverlaps++
		}
	}

	fmt.Printf("\nPart 1: %d", totalContainments)
	fmt.Printf("\nPart 2: %d", totalOverlaps)
}
