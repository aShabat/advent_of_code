package main

import (
	"advent/util"
	"fmt"
)

func main() {
	lines := util.ReadFile("day4/input.txt")

	out1 := 0
	out2 := 0
	for x := range lines {
		for y := range lines[x] {
			for _, d := range directions {
				if check1(lines, x, y, d) {
					out1++
				}
			}
			if check2(lines, x, y) {
				out2++
			}
		}
	}
	fmt.Println(out1, out2)
}

var xmas string = "XMAS"
var directions [][]int = [][]int{
	{1, 0},
	{1, 1},
	{0, 1},
	{-1, 1},
	{-1, 0},
	{-1, -1},
	{0, -1},
	{1, -1},
}

func check1(l []string, x, y int, d []int) bool {
	for i := range 4 {
		if x+i*d[0] < 0 || x+i*d[0] >= len(l) || y+i*d[1] < 0 || y+i*d[1] >= len(l[x+i*d[0]]) || l[x+i*d[0]][y+i*d[1]] != xmas[i] {
			return false
		}
	}
	return true
}

func check2(l []string, x, y int) bool {
	if x < 1 || y < 1 || x >= len(l)-1 || y >= len(l[x])-1 {
		return false
	}

	if l[x][y] != 'A' {
		return false
	}

	if !(l[x-1][y-1] == 'M' && l[x+1][y+1] == 'S') && !(l[x-1][y-1] == 'S' && l[x+1][y+1] == 'M') {
		return false
	}

	if !(l[x-1][y+1] == 'M' && l[x+1][y-1] == 'S') && !(l[x-1][y+1] == 'S' && l[x+1][y-1] == 'M') {
		return false
	}
	return true
}
