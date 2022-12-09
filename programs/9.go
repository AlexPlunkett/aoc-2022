package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Direction string

const (
	Up    Direction = "U"
	Down  Direction = "D"
	Left  Direction = "L"
	Right Direction = "R"
)

type Move struct {
	Direction Direction
	Amount    int
}

type Position struct {
	X int
	Y int
}

func main() {
	contents, _ := os.ReadFile("./inputs/9.txt")
	input := string(contents)

	fmt.Printf("\nPart 1: %d", simulateMovement(input, 2))
	fmt.Printf("\nPart 2: %d", simulateMovement(input, 10))
}

func parseMove(input string) (move Move) {
	parts := regexp.MustCompile(`(U|D|L|R)\s(\d+)`).FindStringSubmatch(input)

	amount, _ := strconv.Atoi(parts[2])

	move.Direction = Direction(parts[1])
	move.Amount = amount

	return move
}

func simulateMovement(moves string, partsCount int) int {
	positions := make([]string, 0, 0)
	positions = append(positions, fmt.Sprint(Position{0, 0}))

	parts := make([]Position, 0, partsCount)
	for i := 0; i < partsCount; i++ {
		parts = append(parts, Position{})
	}

	moveResponses := map[Position]Position{
		// static
		Position{1, 0}:   {0, 0},
		Position{-1, 0}:  {0, 0},
		Position{0, 1}:   {0, 0},
		Position{0, -1}:  {0, 0},
		Position{1, 1}:   {0, 0},
		Position{-1, 1}:  {0, 0},
		Position{1, -1}:  {0, 0},
		Position{-1, -1}: {0, 0},
		// straight
		Position{2, 0}:  {1, 0},
		Position{-2, 0}: {-1, 0},
		Position{0, 2}:  {0, 1},
		Position{0, -2}: {0, -1},
		// diagonal
		Position{1, 2}:   {1, 1},
		Position{-1, 2}:  {-1, 1},
		Position{1, -2}:  {1, -1},
		Position{-1, -2}: {-1, -1},
		Position{2, 1}:   {1, 1},
		Position{-2, 1}:  {-1, 1},
		Position{2, -1}:  {1, -1},
		Position{-2, -1}: {-1, -1},
		Position{2, 2}:   {1, 1},
		Position{-2, 2}:  {-1, 1},
		Position{2, -2}:  {1, -1},
		Position{-2, -2}: {-1, -1},
	}

	for _, row := range strings.Split(moves, "\n") {
		move := parseMove(row)

		for i := 0; i < move.Amount; i++ {
			switch move.Direction {
			case Up:
				parts[len(parts)-1].Y++
			case Down:
				parts[len(parts)-1].Y--
			case Left:
				parts[len(parts)-1].X--
			case Right:
				parts[len(parts)-1].X++
			}

			for x := len(parts) - 2; x >= 0; x-- {
				offset := Position{parts[x+1].X - parts[x].X, parts[x+1].Y - parts[x].Y}
				response := moveResponses[offset]
				parts[x].X += response.X
				parts[x].Y += response.Y
			}

			positionTracked := false

			for _, position := range positions {
				if position == fmt.Sprint(parts[0]) {
					positionTracked = true
					break
				}
			}

			if !positionTracked {
				positions = append(positions, fmt.Sprint(parts[0]))
			}
		}
	}

	return len(positions)
}
