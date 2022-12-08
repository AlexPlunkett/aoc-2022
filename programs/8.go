package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	contents, _ := os.ReadFile("./inputs/8.txt")
	input := string(contents)

	rows := strings.Split(input, "\n")

	grid := make(map[string]int)
	deletion := make([]string, 0, 0)

	for y := 0; y < len(rows); y++ {
		for x := 0; x < len(rows[y]); x++ {
			height, _ := strconv.Atoi(rows[y][x : x+1])
			grid[strconv.Itoa(x)+":"+strconv.Itoa(y)] = height
		}
	}

	scores := make([]int, 0, 0)

	for y := 1; y < len(rows)-1; y++ {
		for x := 1; x < len(rows[y])-1; x++ {
			height, _ := strconv.Atoi(rows[y][x : x+1])
			visibleSides := 4
			sideScores := []int{0, 0, 0, 0}
			// left
			for n := x - 1; n >= 0; n-- {
				sideScores[0]++
				if grid[strconv.Itoa(n)+":"+strconv.Itoa(y)] >= height {
					visibleSides--
					break
				}
			}
			// right
			for n := x + 1; n < len(rows[y]); n++ {
				sideScores[1]++
				if grid[strconv.Itoa(n)+":"+strconv.Itoa(y)] >= height {
					visibleSides--
					break
				}
			}
			// top
			for n := y - 1; n >= 0; n-- {
				sideScores[2]++
				if grid[strconv.Itoa(x)+":"+strconv.Itoa(n)] >= height {
					visibleSides--
					break
				}
			}
			// bottom
			for n := y + 1; n < len(rows); n++ {
				sideScores[3]++
				if grid[strconv.Itoa(x)+":"+strconv.Itoa(n)] >= height {
					visibleSides--
					break
				}
			}

			if visibleSides == 0 {
				deletion = append(deletion, strconv.Itoa(x)+":"+strconv.Itoa(y))
			}

			scores = append(scores, sideScores[0]*sideScores[1]*sideScores[2]*sideScores[3])
		}
	}

	for _, value := range deletion {
		delete(grid, value)
	}

	sort.Ints(scores)

	fmt.Printf("\nPart 1: %d", len(grid))
	fmt.Printf("\nPart 2: %d", scores[len(scores)-1])
}
