package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	if len(os.Args[1:]) == 0 {
		programs, _ := os.ReadDir("./programs")
		for day := 1; day <= len(programs); day++ {
			runProgramOnDay(strconv.Itoa(day))
		}
		return
	}

	day := os.Args[1:][0]

	runProgramOnDay(day)
}

func runProgramOnDay(day string) {
	fmt.Printf("Day %s \n", day)

	out, err := exec.Command("go", "run", fmt.Sprintf("./programs/%s.go", day)).Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Output: %s \n", string(out))
}
