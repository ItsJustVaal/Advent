package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func problem1() {
	// P1
	data, _ := os.ReadFile("../inputs/p1input.txt")
	split := strings.Split(string(data), "\r\n")
	sum := 0
	var first, last, tmpNum string

	for _, x := range split {
		for _, y := range x {
			check := string(y)
			_, err := strconv.Atoi(check)
			if err != nil {
				continue
			} else if first == "" {
				first += string(y)
			} else {
				last = string(y)
			}
		}
		if last == "" {
			tmpNum = first + first
		} else {
			tmpNum = first + last
		}
		numToAdd, err := strconv.Atoi(tmpNum)
		if err != nil {
			fmt.Println(err.Error())
		}
		sum += numToAdd
		//
		first = ""
		last = ""
		tmpNum = ""
	}
	fmt.Println(sum)
}
