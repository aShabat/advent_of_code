package main

import (
	"advent/util"
	"fmt"
	"regexp"
	"slices"
)

func main() {
	lines := util.ReadFile("day5/input.txt")
	order := map[int]map[int]struct{}{}
	isOrder := regexp.MustCompile(`\d\d\|\d\d`)
	for {
		line := lines[0]
		lines = lines[1:]
		if isOrder.MatchString(line) {
			nums := util.Split(line, "|")
			left, right := nums[0], nums[1]
			if _, ok := order[left]; !ok {
				order[left] = map[int]struct{}{}
			}
			order[left][right] = struct{}{}
		} else {
			break
		}
	}

	out1 := 0
	out2 := 0
	for _, line := range lines {
		nums := util.Split(line, ",")
		if slices.IsSortedFunc(nums, func(left, right int) int { return checkOrder(left, right, order) }) {
			out1 += nums[(len(nums)-1)/2]
		} else {
			slices.SortFunc(nums, func(left, right int) int { return checkOrder(left, right, order) })
			out2 += nums[(len(nums)-1)/2]
		}
	}
	fmt.Println(out1, out2)
}

func checkOrder(left, right int, order map[int]map[int]struct{}) int {
	if _, ok := order[left][right]; ok {
		return -1
	} else if _, ok := order[right][left]; ok {
		return 1
	}

	return 0
}
