package main

import (
	"advent/util"
	"fmt"
)

func main() {
	lines := util.ReadFile("day25/input.txt")
	locks, keys := [][]int{}, [][]int{}
	for len(lines) > 0 {
		line := lines[0]
		lines = lines[1:]
		if line == "#####" {
			locks = append(locks, make([]int, 5))
			for range 5 {
				line = lines[0]
				lines = lines[1:]
				for i, r := range line {
					if r == '#' {
						locks[len(locks)-1][i]++
					}
				}
			}
			lines = lines[1:]
		} else if line == "....." {
			keys = append(keys, make([]int, 5))
			for range 5 {
				line = lines[0]
				lines = lines[1:]
				for i, r := range line {
					if r == '#' {
						keys[len(keys)-1][i]++
					}
				}
			}
			lines = lines[1:]
		}
	}
	out1 := 0
	for _, lock := range locks {
		for _, key := range keys {
			overlaps := false
			for i := range 5 {
				overlaps = overlaps || lock[i]+key[i] > 5
			}
			if !overlaps {
				out1++
			}
		}
	}
	fmt.Println(out1)
}
