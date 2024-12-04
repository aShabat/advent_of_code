package main

import (
	"advent/util"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	lines := util.ReadFile("day3/input.txt")

	rx := regexp.MustCompile(`don't\(\)|do\(\)|mul\((\d{1,3}),(\d{1,3})\)`)
	sum1 := 0
	sum2 := 0
	line := ""
	for _, l := range lines {
		line += l
	}
	do := true
	for _, match := range rx.FindAllStringSubmatch(line, -1) {
		if match[0] == "do()" {
			do = true
		} else if match[0] == "don't()" {
			do = false
		} else {
			left, err := strconv.Atoi(match[1])
			if err != nil {
				panic(err)
			}
			right, err := strconv.Atoi(match[2])
			if err != nil {
				panic(err)
			}
			sum1 += left * right
			if do {
				sum2 += left * right
			}
		}
	}

	fmt.Println(sum2)
}
