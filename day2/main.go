package main

import (
	"advent/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines, close := util.ReadFile("day2/input.txt")
	close()
	countSafe1 := 0
	countSafe2 := 0
	for _, line := range lines {
		nums := split(line)
		if isSafe(nums) {
			countSafe1++
		}
		if isSafeWithDampener(nums) {
			countSafe2++
		} else {
			fmt.Println(nums)
		}
	}

	fmt.Println(countSafe1, countSafe2)
}

func split(line string) []int {
	nums := []int{}
	for _, n := range strings.Split(line, " ") {
		num, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)
	}

	return nums
}

func isSafe(nums []int) bool {
	diffs := []int{}
	for i := range len(nums) - 1 {
		diffs = append(diffs, nums[i+1]-nums[i])
	}
	var mult int
	if diffs[0] > 0 {
		mult = 1
	} else {
		mult = -1
	}

	for _, diff := range diffs {
		if mult*diff <= 0 || mult*diff > 3 {
			return false
		}
	}

	return true
}

func isSafeWithDampener(nums []int) bool {
	if isSafe(nums[1:]) {
		return true
	}
	diffs := []int{}
	for i := range len(nums) - 1 {
		diffs = append(diffs, nums[i+1]-nums[i])
	}
	dampened := false
	var mult int

	if diffs[0] > 0 {
		mult = 1
	} else {
		mult = -1
	}

	for i := 0; i < len(diffs); i++ {
		diff := diffs[i]
		if mult*diff <= 0 || mult*diff > 3 {
			if !dampened {
				if i == len(diffs)-1 {
					dampened = true
				} else if temp := mult * (diff + diffs[i+1]); temp >= 1 && temp <= 3 {
					dampened = true
					i++
				} else {
					return false
				}
			} else {
				return false
			}
		}
	}

	return true
}
