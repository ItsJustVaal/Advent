package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type KV struct {
	key   int
	value int
}

// This is just part 2, part 1 I already had completed
func problemOne() {
	data, _ := os.ReadFile("../inputs/p1input.txt")
	split := strings.Split(string(data), "\r\n")

	totals := make(map[int]int)
	sum := 0
	for _, x := range split {
		key := len(totals)
		if x != "" {
			add, _ := strconv.Atoi(x)
			sum += add
		} else {
			totals[key] = sum
			sum = 0
		}
	}
	final := mapToSlice(totals)
	sort.SliceStable(final, func(i, j int) bool {
		return final[i].value > final[j].value
	})
	finalSum := final[0].value + final[1].value + final[2].value
	fmt.Println(finalSum)
}

func mapToSlice(input map[int]int) []KV {
	newSlice := make([]KV, len(input))
	i := 0
	for k, v := range input {
		newSlice[i].key = k
		newSlice[i].value = v
		i++
	}
	return newSlice
}
