package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// P1
	valueMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
	}

	data, _ := os.ReadFile("../inputs/p1input.txt")
	split := strings.Split(string(data), "\r\n")
	sum := 0

	for _, x := range split {
		first := 0
		last := 0
		tmpNum := ""
		for n, y := range x {
			fmt.Printf("TMP IS: %s\n", tmpNum)
			fmt.Printf("FIRST FIRST: %d\n", first)
			fmt.Printf("FIRST LAST: %d\n", last)
			for key, value := range valueMap {
				if strings.Contains(tmpNum, key) {
					if first == 0 {
						first = value
						tmpNum = ""
						break
					} else {
						last = value
						tmpNum = ""
						break
					}
				}
			}
			init, ok := valueMap[tmpNum]
			if !ok {
				fmt.Println("INSIDE NOT OK")

				check, ok := valueMap[string(y)]
				if !ok {
					fmt.Println("NOT FOUND")
					tmpNum += string(y)
					fmt.Println(tmpNum)
					if (n + 1) == len(x) {
						for key, value := range valueMap {
							if strings.Contains(tmpNum, key) {
								if first == 0 {
									first = value
									tmpNum = ""
									break
								} else {
									last = value
									tmpNum = ""
									break
								}
							}
						}
					}
				} else {
					fmt.Println("FOUND")
					if first == 0 {
						first = check
						fmt.Printf("FIRST: %d\n", first)
						tmpNum = ""
					} else {
						last = check
						fmt.Printf("LAST: %d\n", last)
					}
				}

			} else {
				fmt.Println("FOUND OFF THE RIP")
				if first == 0 {
					first = init
					fmt.Printf("FIRST: %d\n", first)
				} else {
					last = init
					fmt.Printf("LAST: %d\n", last)
				}

				tmpNum = string(y)
			}
		}
		if last == 0 {
			for key, value := range valueMap {
				fmt.Printf("KEY: %s\n", key)
				if strings.Contains(tmpNum, key) {
					fmt.Println("FOUND")
					if first == 0 {
						first = value
						break
					} else {
						last = value
						break
					}
				} else {
					fmt.Println("NOT FOUND")
					last = first
				}
			}
		}
		finalCheck, ok := valueMap[tmpNum]
		if ok {
			last = finalCheck
		}
		tmp := fmt.Sprintf("%d%d", first, last)
		fmt.Printf("TMP: %s ", tmp)
		final, err := strconv.Atoi(tmp)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("FINAL: %d\n", final)
		sum += final
	}
	fmt.Println(sum)
}

// P1
// for _, y := range x {
// 	check := string(y)
// 	_, err := strconv.Atoi(check)
// 	if err != nil {
// 		continue
// 	} else if first == "" {
// 		first += string(y)
// 	} else {
// 		last = string(y)
// 	}
// }
// if last == "" {
// 	tmpNum = first + first
// } else {
// 	tmpNum = first + last
// }
// numToAdd, err := strconv.Atoi(tmpNum)
// if err != nil {
// 	fmt.Println(err.Error())
// }
// sum += numToAdd
// //
// first = ""
// last = ""
// tmpNum = ""
