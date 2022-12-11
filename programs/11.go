package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const MonkeyRegex = `Monkey (\d+):
  Starting items: (.+)
  Operation: new = old (.+) (.+)
  Test: divisible by (\d+)
    If true: throw to monkey (\d+)
    If false: throw to monkey (\d+)`

type Operator string

const (
	Add      Operator = "+"
	Subtract Operator = "-"
	Multiply Operator = "*"
	Old      Operator = "old"
)

type Operation struct {
	Value    int64
	ByOld    bool
	Operator Operator
}

type Monkey struct {
	ID                  int
	Items               []int64
	Operation           Operation
	TestDivisor         int64
	AssertTrueMonkeyID  int
	AssertFalseMonkeyID int
	InspectionCount     int64
}

func main() {
	contents, _ := os.ReadFile("./inputs/11.txt")
	input := string(contents)

	monkeysPart1 := make([]Monkey, 0, 0)
	monkeysPart2 := make([]Monkey, 0, 0)

	for _, row := range strings.Split(input, "\n\n") {
		monkeysPart1 = append(monkeysPart1, parseMonkey(row))
		monkeysPart2 = append(monkeysPart2, parseMonkey(row))
	}

	fmt.Printf("\nPart 1: %d", playRounds(monkeysPart1, 20, true))
	fmt.Printf("\nPart 2: %d", playRounds(monkeysPart2, 10000, false))
}

func parseMonkey(input string) (monkey Monkey) {
	parsed := regexp.MustCompile(MonkeyRegex).FindStringSubmatch(input)

	ID, _ := strconv.Atoi(parsed[1])
	monkey.ID = ID

	monkey.Items = make([]int64, 0, 0)
	for _, value := range strings.Split(parsed[2], ", ") {
		worryLevel, _ := strconv.Atoi(value)
		monkey.Items = append(monkey.Items, int64(worryLevel))
	}

	operator := Add
	switch parsed[3] {
	case "*":
		operator = Multiply
	case "+":
		operator = Add
	case "-":
		operator = Subtract
	}

	worryLevel, _ := strconv.Atoi(parsed[4])
	monkey.Operation = Operation{
		Value:    int64(worryLevel),
		ByOld:    parsed[4] == "old",
		Operator: operator,
	}

	testDivisor, _ := strconv.Atoi(parsed[5])
	monkey.TestDivisor = int64(testDivisor)

	assetTrueMonkeyId, _ := strconv.Atoi(parsed[6])
	monkey.AssertTrueMonkeyID = assetTrueMonkeyId

	assetFalseMonkeyId, _ := strconv.Atoi(parsed[7])
	monkey.AssertFalseMonkeyID = assetFalseMonkeyId

	return monkey
}

func playRounds(monkeys []Monkey, limit int, part1 bool) int64 {
	supermod := int64(1)
	for _, monkey := range monkeys {
		supermod *= monkey.TestDivisor
	}

	for i := 0; i < limit; i++ {
		for n, _ := range monkeys {
			for x, _ := range monkeys[n].Items {
				if !part1 {
					monkeys[n].Items[x] %= supermod
				}

				value := monkeys[n].Operation.Value
				if monkeys[n].Operation.ByOld {
					value = monkeys[n].Items[x]
				}

				switch monkeys[n].Operation.Operator {
				case Add:
					monkeys[n].Items[x] += value
				case Subtract:
					monkeys[n].Items[x] -= value
				case Multiply:
					monkeys[n].Items[x] *= value
				}

				if part1 {
					monkeys[n].Items[x] = int64(math.Trunc(float64(monkeys[n].Items[x]) / 3))
				}

				nextMonkeyID := 0
				if monkeys[n].Items[x]%monkeys[n].TestDivisor == 0 {
					nextMonkeyID = monkeys[n].AssertTrueMonkeyID
				} else {
					nextMonkeyID = monkeys[n].AssertFalseMonkeyID
				}

				for p, _ := range monkeys {
					if monkeys[p].ID == nextMonkeyID {
						monkeys[p].Items = append(monkeys[p].Items, monkeys[n].Items[x])
						break
					}
				}

				monkeys[n].InspectionCount++
			}

			monkeys[n].Items = make([]int64, 0, 0)
		}
	}

	sort.SliceStable(monkeys, func(i, j int) bool {
		return monkeys[i].InspectionCount > monkeys[j].InspectionCount
	})

	return monkeys[0].InspectionCount * monkeys[1].InspectionCount
}
