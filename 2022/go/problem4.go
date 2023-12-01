package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func problem4() {
	// Problem 4 P1
	// Check if num 1 in pair 2 is gt= num 1 in pair 1 and tl= num 2
	// Check if num 2 in pair 2 is lt= num 2 in pair 1 and gt= num 1

	data, _ := os.ReadFile("../inputs/p4input.txt")
	split := strings.Split(string(data), "\r\n")
	sum := 0
	for _, x := range split {
		commaSplit := strings.Split(x, ",")
		elfOne := strings.Split(commaSplit[0], "-")
		elfTwo := strings.Split(commaSplit[1], "-")
		num1, _ := strconv.Atoi(elfOne[0])
		num2, _ := strconv.Atoi(elfOne[1])
		num3, _ := strconv.Atoi(elfTwo[0])
		num4, _ := strconv.Atoi(elfTwo[1])

		// P1
		if num3 >= num1 && num4 <= num2 {
			sum++
		} else if num1 >= num3 && num2 <= num4 {
			sum++
		}
		// P2
		if num3 <= num2 && num3 >= num1 {
			sum++
		} else if num1 >= num3 && num1 <= num4 {
			sum++
		}

	}
	fmt.Println(sum)
}
