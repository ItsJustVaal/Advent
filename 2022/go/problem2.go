package main

import (
	"fmt"
	"os"
	"strings"
)

func problem2() {
	type guide struct {
		win []string
		tie []string
	}
	// Part One
	winning := [3]string{"A Y", "B Z", "C X"}
	tying := [3]string{"A X", "B Y", "C Z"}
	cheatGuide := guide{
		win: winning[:],
		tie: tying[:],
	}

	input, err := os.ReadFile("../inputs/p2input.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	split := strings.Split(string(input), "\r\n")

	var sum int
	for _, x := range split {
		for _, y := range cheatGuide.win {
			if x == y {
				sum += 6
			}
		}
		for _, y := range cheatGuide.tie {
			if x == y {
				sum += 3
			}
		}
		switch string(x[2]) {
		case "X":
			sum++
		case "Y":
			sum += 2
		case "Z":
			sum += 3
		}

	}

	fmt.Println(sum)

	// Part 2
	sumTwo := 0
	values := make(map[string]int)
	values["A"] = 1
	values["B"] = 2
	values["C"] = 3

	oppositeWins := make(map[string]string)
	oppositeWins["A"] = "B"
	oppositeWins["B"] = "C"
	oppositeWins["C"] = "A"

	oppositeLoses := make(map[string]string)
	oppositeLoses["A"] = "C"
	oppositeLoses["B"] = "A"
	oppositeLoses["C"] = "B"

	for _, x := range split {
		switch string(x[2]) {
		case "X":
			sumTwo += values[oppositeLoses[string(x[0])]]
		case "Y":
			sumTwo += 3 + (values[string(x[0])])
		case "Z":
			sumTwo += 6 + (values[oppositeWins[string(x[0])]])
		}
	}
	fmt.Println(sumTwo)
}
