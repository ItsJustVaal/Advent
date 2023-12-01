package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func problem5() {
	// 		   [J]         [B]     [T]
	// 		   [M] [L]     [Q] [L] [R]
	// 		   [G] [Q]     [W] [S] [B] [L]
	// [D]     [D] [T]     [M] [G] [V] [P]
	// [T]     [N] [N] [N] [D] [J] [G] [N]
	// [W] [H] [H] [S] [C] [N] [R] [W] [D]
	// [N] [P] [P] [W] [H] [H] [B] [N] [G]
	// [L] [C] [W] [C] [P] [T] [M] [Z] [W]
	// 1   2   3   4   5   6   7   8   9
	boxMap := [][]string{
		{"L", "N", "W", "T", "D"},
		{"C", "P", "H"},
		{"W", "P", "H", "N", "D", "G", "M", "J"},
		{"C", "W", "S", "N", "T", "Q", "L"},
		{"P", "H", "C", "N"},
		{"T", "H", "N", "D", "M", "W", "Q", "B"},
		{"M", "B", "R", "J", "G", "S", "L"},
		{"Z", "N", "W", "G", "V", "B", "R", "T"},
		{"W", "G", "D", "N", "P", "L"},
	}

	data, _ := os.ReadFile("../inputs/p5input.txt")
	split := strings.Split(string(data), "\r\n")
	for i, x := range split {
		queue := make([]string, 0)
		split[i] = strings.Replace(x, "move ", "", -1)
		split[i] = strings.Replace(split[i], " from", "", -1)
		split[i] = strings.Replace(split[i], " to", "", -1)
		checks := strings.Split(split[i], " ")
		toMove, _ := strconv.Atoi(checks[0])
		moveFrom, _ := strconv.Atoi(checks[1])
		moveTo, _ := strconv.Atoi(checks[2])
		fmt.Printf("Moving From: %d, Moving To: %d, Num: %d\n", moveFrom, moveTo, toMove)
		// P1
		queue, newBoxStack := enqueue(boxMap[moveFrom-1], queue, toMove)
		fmt.Println("post enqueue return")
		boxMap[moveFrom-1] = newBoxStack
		fmt.Println("adding")
		changedBoxStack := addToStack(boxMap[moveTo-1], queue)
		boxMap[moveTo-1] = changedBoxStack

		// P2
		queue, newBoxStack := enqP2(boxMap[moveFrom-1], queue, toMove)
		boxMap[moveFrom-1] = newBoxStack
		changedBoxStack := addToStack(boxMap[moveTo-1], queue)
		boxMap[moveTo-1] = changedBoxStack
		for _, j := range boxMap {
			fmt.Println(j)
		}

	}
	fmt.Println("FIN")
	for _, x := range boxMap {
		fmt.Println(x)
	}

}

// P1
func enqueue(boxStack, queue []string, numberToMove int) ([]string, []string) {
	for x := 0; x < numberToMove; x++ {
		fmt.Println("enqueue")
		lastItem := boxStack[len(boxStack)-1]
		queue = append(queue, lastItem)
		boxStack = removeFromStack(len(boxStack)-1, boxStack)
		fmt.Println("return")
	}
	return queue, boxStack
}

func dequeue(queue []string) []string {
	fmt.Println("dequeue")
	if len(queue) == 0 {
		return []string{}
	}
	return queue[1:]
}

func removeFromStack(index int, boxStack []string) []string {
	if len(boxStack) == 0 || index == 0 {
		fmt.Println("returning")
		return boxStack
	}
	fmt.Println("RETURNING")
	return boxStack[:index-1]
}

// P2
func addToStack(boxStack, queue []string) []string {
	boxStack = append(boxStack, queue...)
	return boxStack
}

func enqP2(boxStack, queue []string, numberToMove int) ([]string, []string) {
	num := len(boxStack) - numberToMove
	queue = boxStack[num:]
	boxStack = boxStack[:num]
	return queue, boxStack
}
