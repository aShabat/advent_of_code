package main

import (
	"advent/util"
	"fmt"
	"strings"
)

func main() {
	lines := util.ReadFile("day19/input.txt")
	towels := strings.Split(lines[0], ", ")
	countPossible := countPossibleGen(towels)
	lines = lines[2:]
	out1 := 0
	out2 := 0
	for _, line := range lines {
		if countPossible(line) != 0 {
			out1++
		}
		out2 += countPossible(line)
	}
	fmt.Println(out1, out2)
}

func countPossibleGen(towels []string) func(string) int {
	countPossibleMemo := map[string]int{}
	countPossibleMemo[""] = 1
	var countPossible func(string) int
	countPossible = func(design string) int {
		if out, ok := countPossibleMemo[design]; ok {
			return out
		}
		countPossibleMemo[design] = 0
		for _, towel := range towels {
			if d, ok := strings.CutPrefix(design, towel); ok {
				countPossibleMemo[design] += countPossible(d)
			}
		}
		return countPossibleMemo[design]
	}

	return countPossible
}
