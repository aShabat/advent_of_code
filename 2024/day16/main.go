package main

import (
	"advent/util"
	"fmt"
	"slices"
)

func main() {
	lines := util.ReadFile("day16/input.txt")
	maze := make([][]int, len(lines))
	var rRow, rCol, fRow, fCol int
	for row, line := range lines {
		maze[row] = make([]int, len(line))
		for col, l := range line {
			if l == '#' {
				maze[row][col] = -1
			}
			if l == 'S' {
				rRow, rCol = row, col
			}
			if l == 'E' {
				fRow, fCol = row, col
			}
		}
	}
	fmt.Println(solveMaze(maze, rRow, rCol, fRow, fCol))
}

var directions = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func solveMaze(maze [][]int, rRow, rCol, fRow, fCol int) (int, int) {
	dists := make([][][]int, len(maze))
	goodAncestors := make([][][][][]int, len(maze))
	counted := make([][]bool, len(maze))
	points := [][]int{}
	for row := range maze {
		dists[row] = make([][]int, len(maze[row]))
		goodAncestors[row] = make([][][][]int, len(maze[row]))
		counted[row] = make([]bool, len(maze[row]))
		for col := range dists[row] {
			dists[row][col] = make([]int, 4)
			goodAncestors[row][col] = make([][][]int, 4)
			for dir := range 4 {
				dists[row][col][dir] = -1
				points = append(points, []int{row, col, dir})
			}
		}
	}
	dists[rRow][rCol][0] = 0
	sort(dists, points)
	var score int
	for len(points) > 0 {
		fmt.Println(len(points))
		now := points[0]
		points = points[1:]
		if dists[now[0]][now[1]][now[2]] == -1 {
			break
		}
		if now[0] == fRow && now[1] == fCol {
			score = dists[now[0]][now[1]][now[2]]
		}
		nexts := [][]int{
			{now[0] + directions[now[2]][0], now[1] + directions[now[2]][1], now[2], 1},
			{now[0], now[1], (now[2] + 1) % 4, 1000},
			{now[0], now[1], (now[2] + 2) % 4, 2000},
			{now[0], now[1], (now[2] + 3) % 4, 1000},
		}
		for _, next := range nexts {
			if maze[next[0]][next[1]] == -1 {
				continue
			}
			if nd := dists[next[0]][next[1]][next[2]]; nd == -1 || nd > dists[now[0]][now[1]][now[2]]+next[3] {
				dists[next[0]][next[1]][next[2]] = dists[now[0]][now[1]][now[2]] + next[3]
				goodAncestors[next[0]][next[1]][next[2]] = [][]int{{now[0], now[1], now[2]}}
				sort(dists, points)
			} else if nd == dists[now[0]][now[1]][now[2]]+next[3] {
				goodAncestors[next[0]][next[1]][next[2]] = append(goodAncestors[next[0]][next[1]][next[2]], []int{now[0], now[1], now[2]})
			}
		}
	}

	set := [][]int{}
	countSeats := 0
	for dir := range 4 {
		if dists[fRow][fCol][dir] == score {
			set = append(set, []int{fRow, fCol, dir})
		}
	}

	for len(set) > 0 {
		now := set[0]
		set = set[1:]
		if !counted[now[0]][now[1]] {
			countSeats++
			counted[now[0]][now[1]] = true
		}
		for _, a := range goodAncestors[now[0]][now[1]][now[2]] {
			set = append(set, a)
		}
	}
	return score, countSeats
}

func sort(dists [][][]int, points [][]int) {
	slices.SortFunc(points, func(left, right []int) int {
		ld := dists[left[0]][left[1]][left[2]]
		rd := dists[right[0]][right[1]][right[2]]
		if ld == -1 && rd == -1 {
			return 0
		} else if ld == -1 {
			return 1
		} else if rd == -1 {
			return -1
		} else {
			return ld - rd
		}
	})
}
