package main

import (
	"advent/util"
	"fmt"
)

func main() {
	lines := util.ReadFile("day12/input.txt")
	garden := make([][]byte, len(lines))
	for i, line := range lines {
		garden[i] = []byte(line)
	}

	visited := make([][]bool, len(garden))
	for i := range garden {
		visited[i] = make([]bool, len(garden[i]))
	}
	out1 := 0
	out2 := 0
	queue := [][]int{}
	directions := [][]int{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}
	for row := range garden {
		for col := range garden[row] {
			if !visited[row][col] {
				area, perimeter, sides := 0, 0, 0
				queue = [][]int{{row, col}}
				visited[row][col] = true
				for len(queue) > 0 {
					area++
					now := queue[0]
					queue = queue[1:]
					for d, dir := range directions {
						next := []int{now[0] + dir[0], now[1] + dir[1]}
						if !util.InBounds(garden, next[0], next[1]) {
							perimeter++

							shiftDir := directions[(d+1)%4]
							nowShift := []int{now[0] + shiftDir[0], now[1] + shiftDir[1]}
							if !util.InBounds(garden, nowShift[0], nowShift[1]) ||
								garden[now[0]][now[1]] != garden[nowShift[0]][nowShift[1]] {
								sides++
							}
						} else if garden[now[0]][now[1]] != garden[next[0]][next[1]] {
							perimeter++

							shiftDir := directions[(d+1)%4]
							nowShift := []int{now[0] + shiftDir[0], now[1] + shiftDir[1]}
							nextShift := []int{next[0] + shiftDir[0], next[1] + shiftDir[1]}
							if !util.InBounds(garden, nowShift[0], nowShift[1]) ||
								garden[now[0]][now[1]] == garden[nextShift[0]][nextShift[1]] ||
								garden[now[0]][now[1]] != garden[nowShift[0]][nowShift[1]] {
								sides++
							}
						} else if !visited[next[0]][next[1]] {
							visited[next[0]][next[1]] = true
							queue = append(queue, next)
						}
					}
				}
				out1 += area * perimeter
				out2 += area * sides
			}
		}
	}
	fmt.Println(out1, out2)
}
