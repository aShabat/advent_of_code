package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	var out int64 = 0
	left := []int64{}
	right := map[int64]int64{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, "   ")
		if len(nums) != 2 {
			panic("wrong input")
		}
		a, err := strconv.ParseInt(nums[0], 10, 64)
		if err != nil {
			panic(err)
		}
		left = append(left, a)

		b, err := strconv.ParseInt(nums[1], 10, 64)
		if err != nil {
			panic(err)
		}
		if _, ok := right[b]; !ok {
			right[b] = 1
		} else {
			right[b]++
		}
	}
	for _, l := range left {
		out += l * right[l]
	}

	fmt.Println(out)
}
