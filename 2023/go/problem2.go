package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func problem2() {
	// P2 - Part 1
	// Go through each line of the file and check if
	// the current value of the current color is gt=
	// the limits, sum++ else continue
	// This is gross but it works
	// Part 2
	// Multi the max values of each color in each game

	data, _ := os.ReadFile("../inputs/p2input.txt")
	split := strings.Split(string(data), "\r\n")
	// const red = 12
	// const blue = 14
	// const green = 13
	sum := 0
	// flag := 0
	lines := make([]string, 0)
	for _, x := range split {
		y := strings.Split(x, ": ")[1]
		lines = append(lines, y)
	}
	for _, p := range lines {
		var maxRed, maxGreen, maxBlue int
		h := strings.Split(p, " ")
		for w, l := range h {
			if strings.Contains(l, "red") {
				tmp, _ := strconv.Atoi(h[w-1])
				if tmp > maxRed {
					maxRed = tmp
				}
				// if tmp > red {
				// 	fmt.Println("BROKE ON RED")
				// 	// flag = 1
				// 	break
				// }
			} else if strings.Contains(l, "green") {
				tmp, _ := strconv.Atoi(h[w-1])
				if tmp > maxGreen {
					maxGreen = tmp
				}
				// if tmp > green {
				// 	fmt.Println("BROKE ON GREEN")
				// 	// flag = 1
				// 	break
				// }
			} else if strings.Contains(l, "blue") {
				tmp, _ := strconv.Atoi(h[w-1])
				if tmp > maxBlue {
					maxBlue = tmp
				}
				// if tmp > blue {
				// 	fmt.Println("BROKE ON BLUE")
				// 	// flag = 1
				// 	break
				// }
			}

			// flag = 0
		}
		// P2
		fmt.Printf("MAX RED: %d, MAX GREEN %d, MAX BLUE: %d\n", maxRed, maxGreen, maxBlue)
		sum += (maxRed * maxBlue * maxGreen)

		// P1
		// if flag != 1 {
		// 	fmt.Printf("ADDING %d\n", i+1)
		// 	sum += i + 1
		// }
	}
	fmt.Println(sum)
}
