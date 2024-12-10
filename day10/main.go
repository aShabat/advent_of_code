package main

import (
	"advent/util"
	"fmt"
)

func main() {
	lines := util.ReadFile("day10/input.txt")
	nums := make([][]int, len(lines))
	for i, line := range lines {
		nums[i] = util.Split(line, "")
	}
	out1 := 0
	out2 := 0
	for row := range nums {
		for col := range nums[row] {
			out1 += countTrailHeads(nums, row, col)
			if nums[row][col] == 0 {
				out2 += rateTrailHeads(nums, row, col)
			}
		}
	}
	fmt.Println(out1, out2)
}

var directions = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func countTrailHeadsHelper(nums [][]int, visited [][]bool, row, col int) int {
	if !util.InBounds(nums, row, col) || visited[row][col] {
		return 0
	}
	visited[row][col] = true
	if nums[row][col] == 9 {
		return 1
	}
	count := 0
	for _, direction := range directions {
		if util.InBounds(nums, row+direction[0], col+direction[1]) && nums[row][col]+1 == nums[row+direction[0]][col+direction[1]] {
			count += countTrailHeadsHelper(nums, visited, row+direction[0], col+direction[1])
		}
	}
	return count
}

func countTrailHeads(nums [][]int, row, col int) int {
	if !util.InBounds(nums, row, col) || nums[row][col] != 0 {
		return 0
	}
	visited := make([][]bool, len(nums))
	for i := range nums {
		visited[i] = make([]bool, len(nums[i]))
	}
	return countTrailHeadsHelper(nums, visited, row, col)
}

func rateTrailHeads(nums [][]int, row, col int) int {
	if !util.InBounds(nums, row, col) {
		return 0
	}
	if nums[row][col] == 9 {
		return 1
	}
	count := 0
	for _, direction := range directions {
		if util.InBounds(nums, row+direction[0], col+direction[1]) && nums[row][col]+1 == nums[row+direction[0]][col+direction[1]] {
			count += rateTrailHeads(nums, row+direction[0], col+direction[1])
		}
	}
	return count
}
