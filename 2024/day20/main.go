package main

import (
	"advent/util"
	"fmt"
)

func main() {
	lines := util.ReadFile("day20/input.txt")
	race := make([][]int, len(lines))
	distances := make([][][]int, len(lines))
	var sR, sC, eR, eC int
	for row, line := range lines {
		race[row] = make([]int, len(line))
		distances[row] = make([][]int, len(line))
		for col, char := range line {
			if char == '#' {
				race[row][col] = -1
			} else {
				race[row][col] = 0
			}
			if char == 'S' {
				sR, sC = row, col
			}
			if char == 'E' {
				eR, eC = row, col
			}
			distances[row][col] = []int{-1, -1}
		}
	}
	distances[sR][sC][0] = 0
	distances[eR][eC][1] = 0
	stack := [][]int{{sR, sC}}
	for len(stack) > 0 {
		now := stack[0]
		stack = stack[1:]
		for _, dir := range dirs {
			next := []int{now[0] + dir[0], now[1] + dir[1]}
			if race[next[0]][next[1]] != -1 && distances[next[0]][next[1]][0] == -1 {
				distances[next[0]][next[1]][0] = distances[now[0]][now[1]][0] + 1
				stack = append(stack, next)
			}
		}
	}
	stack = [][]int{{eR, eC}}
	for len(stack) > 0 {
		now := stack[0]
		stack = stack[1:]
		for _, dir := range dirs {
			next := []int{now[0] + dir[0], now[1] + dir[1]}
			if race[next[0]][next[1]] != -1 && distances[next[0]][next[1]][1] == -1 {
				distances[next[0]][next[1]][1] = distances[now[0]][now[1]][1] + 1
				stack = append(stack, next)
			}
		}
	}
	out1 := 0
	out2 := 0
	time := distances[eR][eC][0]
	for csR := range race {
		for csC := range race[csR] {
			for cfR := range race {
				for cfC := range race[cfR] {
					if abs(csR-cfR)+abs(csC-cfC) == 2 && distances[csR][csC][0]+distances[cfR][cfC][1]+2 <= time-100 &&
						distances[csR][csC][0] != -1 && distances[cfR][cfC][1] != -1 {
						out1++
					}
					if abs(csR-cfR)+abs(csC-cfC) <= 20 &&
						distances[csR][csC][0]+distances[cfR][cfC][1]+abs(csR-cfR)+abs(csC-cfC) <= time-100 &&
						distances[csR][csC][0] != -1 && distances[cfR][cfC][1] != -1 {
						out2++
					}
				}
			}
		}
	}
	fmt.Println(out1, out2)
}

var dirs = [][]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}
