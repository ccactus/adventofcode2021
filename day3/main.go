package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type CommonCounter struct {
	HighCount int
	LowCount  int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// var counter []CommonCounter // We'll need this to be the exact length
	file, err := os.Open("input.txt") // Reads full file in one go.
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var diag []string
	for scanner.Scan() {
		diag = append(diag, scanner.Text())
	}

	fmt.Printf("Number of entries to consider = %d\n", len(diag[0]))
	counter := make([]CommonCounter, len(diag[0]))
	for _, c := range diag {
		for i, r := range c {
			if r == rune('1') {
				counter[i].HighCount += 1
			} else {
				counter[i].LowCount += 1
			}
		}
	}

	var gammaS, epsilonS string
	for _, b := range counter {
		if b.HighCount < b.LowCount {
			gammaS = gammaS + "0"
			epsilonS = epsilonS + "1"
		} else {
			gammaS = gammaS + "1"
			epsilonS = epsilonS + "0"
		}
	}
	gamma, _ := strconv.ParseUint(gammaS, 2, 16)
	epsilon, _ := strconv.ParseUint(epsilonS, 2, 16)
	fmt.Printf("The power consumption is %d", gamma*epsilon)
}
