package main

/*
The first order of business is to figure out how quickly
the depth icnreases, just so you know what you're dealing
with - you never know if the keys will get carried into
deeper water by an ocean current or a fish or something.
To do this, count hte number of times a depth measurement
increasaes from previous measurement (there's no
measurement before the fist one)
*/

// Read input
// Duplicate and shift left
// parallel run diff to create increase/decrease notes

import (
	"fmt"
)

func checkIncrease(a int, b int) int {
	if a < b {
		return 1
	}
	return 0
}

func main() {
	input1, input2 := Input[:len(Input)-1], Input[1:]
	fmt.Printf("Input1: %d, Input2: %d\n", len(input1), len(input2))
	var count int
	for i, _ := range input1 {
		count = count + checkIncrease(input1[i], input2[i])
	}
	fmt.Println(count)
}
