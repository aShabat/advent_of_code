package util

import (
	"strconv"
	"strings"
)

func Split(line, separator string) []int {
	stringNums := strings.Split(line, separator)
	nums := make([]int, len(stringNums))
	for i, sn := range stringNums {
		num, err := strconv.Atoi(sn)
		if err != nil {
			panic(err)
		}
		nums[i] = num
	}

	return nums
}
