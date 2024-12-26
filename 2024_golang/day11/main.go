package main

import (
	"advent/util"
	"fmt"
)

func main() {
	lines := util.ReadFile("day11/input.txt")
	nums := util.Split(lines[0], " ")

	multiBlinkCount := multiBlinkCountGen()
	out1, out2 := 0, 0
	for _, num := range nums {
		out1 += multiBlinkCount(num, 25)
		out2 += multiBlinkCount(num, 75)
	}
	fmt.Println(out1, out2)
}

func blink(num int) (int, int, bool) {
	if num == 0 {
		return 1, 0, false
	}
	big, small := 100, 10
	for num >= big {
		big *= 100
		small *= 10
	}
	if num/(big/10) == 0 {
		return num * 2024, 0, false
	}
	return num / small, num % small, true
}

func multiBlinkCountGen() func(int, int) int {
	memory := make([]map[int]int, 100)
	for i := range memory {
		memory[i] = map[int]int{}
	}

	var multiBlinkCount func(int, int) int
	multiBlinkCount = func(num, repeat int) int {
		if count, ok := memory[repeat][num]; ok {
			return count
		}
		var count int
		if repeat == 0 {
			count = 1
		} else {
			left, right, two := blink(num)
			count = multiBlinkCount(left, repeat-1)
			if two {
				count += multiBlinkCount(right, repeat-1)
			}
		}
		memory[repeat][num] = count
		return count
	}

	return multiBlinkCount
}
