package main

import (
	"advent/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := util.ReadFile("day7/input.txt")

	out1, out2 := 0, 0
	for _, line := range lines {
		out1 += isSolvable1(line)
		out2 += isSolvable2(line)
	}

	fmt.Println(out2)
}

func isSolvable1(line string) int {
	goalString, line, _ := strings.Cut(line, ": ")
	goal, _ := strconv.Atoi(goalString)
	nums := util.Split(line, " ")

	for bits := range pow(2, len(nums)-1) {
		sum := nums[0]
		for i := range len(nums) - 1 {
			i++
			if bits%2 == 0 {
				sum += nums[i]
			} else {
				sum *= nums[i]
			}
			bits /= 2
		}

		if sum == goal {
			return goal
		}
	}
	return 0
}

func isSolvable2(line string) int {
	goalString, line, _ := strings.Cut(line, ": ")
	goal, _ := strconv.Atoi(goalString)
	nums := util.Split(line, " ")

	for bits := range pow(3, len(nums)-1) {
		if eval2(nums, bits) == goal {
			return goal
		}
	}
	return 0
}

func eval2(nums []int, bits int) int {
	sum := nums[0]
	for i := range len(nums) - 1 {
		i++
		if bits%3 == 0 {
			sum += nums[i]
		} else if bits%3 == 1 {
			sum *= nums[i]
		} else {
			sum, _ = strconv.Atoi(strconv.Itoa(sum) + strconv.Itoa(nums[i]))
		}
		bits /= 3
	}
	return sum
}

func pow(p, n int) int {
	if n == 0 {
		return 1
	} else {
		return p * pow(p, n-1)
	}
}
