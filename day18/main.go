package main

import (
	"advent/util"
	"fmt"
	"slices"
)

func main() {
	lines := util.ReadFile("day18/input.txt")
	blocks := make([][2]int, len(lines))
	for i, line := range lines {
		copy(blocks[i][:], util.Split(line, ",")[:2])
	}
	spaceWidth := 71
	space := make([][]int, spaceWidth)
	distances := make([][]int, spaceWidth)
	unvisited := make([][]int, 0, 4900)
	for row := range spaceWidth {
		space[row] = make([]int, spaceWidth)
		distances[row] = make([]int, spaceWidth)
		for col := range spaceWidth {
			distances[row][col] = -1
			unvisited = append(unvisited, []int{row, col})
		}
	}
	for _, block := range blocks[:1024] {
		space[block[0]][block[1]] = -1
	}
	distances[0][0] = 0
	sortUnvisited := func() {
		slices.SortFunc(unvisited, func(left, right []int) int {
			dl, dr := distances[left[0]][left[1]], distances[right[0]][right[1]]
			if dr == -1 {
				return -1
			} else if dl == -1 {
				return 1
			} else {
				return dl - dr
			}
		})
	}
	directions := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	for len(unvisited) > 0 {
		sortUnvisited()
		now := unvisited[0]
		unvisited = unvisited[1:]
		if distances[now[0]][now[1]] == -1 {
			break
		}
		for _, dir := range directions {
			next := []int{now[0] + dir[0], now[1] + dir[1]}
			if util.InBounds(distances, next[0], next[1]) &&
				space[next[0]][next[1]] != -1 &&
				(distances[next[0]][next[1]] > distances[now[0]][now[1]]+1 || distances[next[0]][next[1]] == -1) {
				distances[next[0]][next[1]] = distances[now[0]][now[1]] + 1
			}
		}
	}
	out1 := distances[spaceWidth-1][spaceWidth-1]
	fmt.Println(out1)

	directions = append(directions, [][]int{
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
	}...)

	blockSets := [][][2]int{}
	isWall := [][]bool{}
	for _, block := range blocks {
		near := []int{}
		for _, dir := range directions {
			for i := range blockSets {
				if slices.Contains(blockSets[i], [2]int{block[0] + dir[0], block[1] + dir[1]}) && !slices.Contains(near, i) {
					near = append(near, i)
				}
			}
		}
		var i int
		if len(near) == 0 {
			i = len(blockSets)
			blockSets = append(blockSets, [][2]int{block})
			isWall = append(isWall, []bool{block[0] == 0 || block[1] == spaceWidth-1, block[0] == spaceWidth-1 || block[1] == 0})
		} else if len(near) == 1 {
			i = near[0]
			blockSets[i] = append(blockSets[i], block)
			isWall[i][0] = isWall[i][0] || block[0] == 0 || block[1] == spaceWidth-1
			isWall[i][1] = isWall[i][1] || block[0] == spaceWidth-1 || block[1] == 0
		} else {
			i = near[0]
			for _, j := range near[1:] {
				blockSets[i] = append(blockSets[i], blockSets[j]...)
				isWall[i][0] = isWall[i][0] || isWall[j][0]
				isWall[i][1] = isWall[i][1] || isWall[j][1]
				blockSets[j] = [][2]int{}
			}
			blockSets[i] = append(blockSets[i], block)
			isWall[i][0] = isWall[i][0] || block[0] == 0 || block[1] == spaceWidth-1
			isWall[i][1] = isWall[i][1] || block[0] == spaceWidth-1 || block[1] == 0
		}
		if isWall[i][0] && isWall[i][1] {
			fmt.Println(block)
			break
		}
		space[block[0]][block[1]] = -1
		// printSpace(space)
		// fmt.Scanln()
	}
}

func printSpace(space [][]int) {
	for row := range space {
		for cell := range space[row] {
			if space[row][cell] == -1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
