package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const (
	CommandCd = "cd"
	CommandLs = "ls"
)

func main() {
	contents, _ := os.ReadFile("./inputs/7.txt")
	input := string(contents)

	var cmd, arg string
	cwd := make([]string, 0, 0)
	directories := make(map[string]int)
	checkedFiles := make([]string, 0, 0)

	for _, row := range strings.Split(input, "\n") {
		if row[:1] == "$" {
			cmd, arg = parseCommand(row)

			if cmd == CommandCd {
				if arg == ".." {
					cwd = cwd[:len(cwd)-1]
				} else if arg != "/" {
					cwd = append(cwd, arg)
				}
				continue
			}
		}

		if cmd == CommandLs {
			directory := "/" + strings.Join(cwd[:], "/")

			if _, ok := directories[directory]; !ok {
				directories[directory] = 0
			}

			if ok, size, name := readFileInfo(row); ok {
				file := directory + "/" + name
				hasChecked := false

				for _, v := range checkedFiles {
					if v == file {
						hasChecked = true
					}
				}

				if !hasChecked {
					for i, _ := range cwd {
						directories["/"+strings.Join(cwd[:i+1], "/")] += size
					}
					directories["/"] += size
					checkedFiles = append(checkedFiles, file)
				}
			}
		}
	}

	var totalSize int
	minimumSizes := make([]int, 0, 0)

	for _, value := range directories {
		if value <= 100000 {
			totalSize += value
		}

		if value+(70000000-directories["/"]) >= 30000000 {
			minimumSizes = append(minimumSizes, value)
		}
	}

	sort.Ints(minimumSizes)

	fmt.Printf("\nPart 1: %d", totalSize)
	fmt.Printf("\nPart 2: %d", minimumSizes[0])
}

func parseCommand(input string) (cmd, arg string) {
	parsed := regexp.MustCompile(`\$\s(\w+)\s?(.+)?`).FindStringSubmatch(input)

	cmd = parsed[1]

	if len(parsed) == 3 {
		arg = parsed[2]
	}

	return cmd, arg
}

func readFileInfo(input string) (ok bool, size int, name string) {
	parsed := regexp.MustCompile(`(\d+)\s(.+)`).FindStringSubmatch(input)

	if len(parsed) == 3 {
		ok = true
		bytes, _ := strconv.Atoi(parsed[1])
		size = bytes
		name = parsed[2]
	} else {
		ok = false
	}

	return ok, size, name
}
