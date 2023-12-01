package main

import (
	"fmt"
	"os"
	"strings"
)

func problem3() {
	// problem 3
	// split the string in half
	// find common letters
	// for each common letter add value from map to total
	// return total
	total := 0
	values := make(map[rune]int)
	lcCurr := 'a'
	upCurr := 'A'
	for x := 1; x <= 26; x++ {
		values[lcCurr] = x
		lcCurr++
	}
	for x := 27; x <= 52; x++ {
		values[upCurr] = x
		upCurr++
	}

	input, err := os.ReadFile("../inputs/p3input.txt")
	if err != nil {
		fmt.Println(err.Error())
	}

	split := strings.Split(string(input), "\r\n")
	for _, x := range split {
		left := x[:len(x)/2]
		right := x[len(x)/2:]

		for _, x := range left {
			if strings.ContainsRune(right, x) {
				val, ok := values[x]
				if !ok {
					continue
				}
				total += val
				break
			}
		}
	}
	fmt.Println(total)

	// part two
	totalTwo := 0
	for x := 0; x < len(split); x += 3 {
		var group []string
		group = append(group, split[x])
		group = append(group, split[x+1])
		group = append(group, split[x+2])
		for _, y := range group[0] {
			if strings.ContainsRune(group[1], y) && strings.ContainsRune(group[2], y) {
				totalTwo += values[y]
				break
			}
		}
	}
	fmt.Println(totalTwo)
}
