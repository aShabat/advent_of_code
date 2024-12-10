package main

import (
	"advent/util"
	"fmt"
)

func main() {
	lines := util.ReadFile("day9/input.txt")
	line := lines[0]
	nums := make([]int, len(line))
	for i, l := range line {
		nums[i] = int(l - '0')
	}
	var out1 uint64
	head, hTaken := 0, 0
	tail, tTaken := len(nums)-1, 0
	if tail%2 == 1 {
		tail--
	}
	for i := 0; ; i++ {
		for nums[head] == 0 {
			head++
		}
		for nums[tail] == 0 {
			tail -= 2
		}
		if head == tail && hTaken+tTaken == nums[head] || tail < head {
			break
		}
		if head%2 == 0 {
			out1 += uint64((head / 2) * i)
			hTaken++
			if hTaken == nums[head] {
				head++
				hTaken = 0
			}
		} else {
			out1 += uint64((tail / 2) * i)
			tTaken++
			if tTaken == nums[tail] {
				tail -= 2
				tTaken = 0
			}
			hTaken++
			if hTaken == nums[head] {
				head++
				hTaken = 0
			}
		}
	}
	fmt.Println(out1)

	numBlocks := [][]int{}
	wsBlocks := [][]int{}
	pos := 0
	for i := range nums {
		if i%2 == 0 {
			numBlocks = append(numBlocks, []int{nums[i], pos})
		} else {
			wsBlocks = append(wsBlocks, []int{nums[i], pos})
		}
		pos += nums[i]
	}
	for i := len(numBlocks) - 1; i >= 0; i-- {
		for j := range wsBlocks {
			if j >= i {
				break
			}
			if numBlocks[i][0] <= wsBlocks[j][0] {
				numBlocks[i][1] = wsBlocks[j][1]
				wsBlocks[j][0] -= numBlocks[i][0]
				wsBlocks[j][1] += numBlocks[i][0]
				break
			}
		}
	}
	var out2 uint
	for i := range numBlocks {
		for j := range numBlocks[i][0] {
			out2 += uint(i * (numBlocks[i][1] + j))
		}
	}
	fmt.Println(out2)
}
