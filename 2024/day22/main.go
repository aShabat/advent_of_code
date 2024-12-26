package main

import (
	"advent/util"
	"fmt"
	"strconv"
)

func main() {
	lines := util.ReadFile("day22/input.txt")
	out1 := 0
	prices := make([][]int, len(lines))
	for i, line := range lines {
		secret, _ := strconv.Atoi(line)
		secret, prices[i] = update(secret, 2000)
		out1 += secret
	}
	fmt.Println(out1)
	caches := make([][][][][]int, len(prices))
	for i := range caches {
		caches[i] = genInstructionCache(prices[i])
	}
	out2 := 0
	for i := range 19 * 19 * 19 * 19 {
		a, i := i%19, i/19
		b, i := i%19, i/19
		c, i := i%19, i/19
		d := i % 19
		totalSell := 0
		for j := range prices {
			if caches[j][a][b][c][d] != -1 {
				totalSell += caches[j][a][b][c][d]
			}
		}
		if totalSell > out2 {
			out2 = totalSell
		}
	}
	fmt.Println(out2)
}

func mixAndPrune(a, b int) int {
	return (a ^ b) % 16777216
}

func update(secret, times int) (int, []int) {
	prices := []int{secret % 10}
	for range times {
		secret = mixAndPrune(secret, secret*64)
		secret = mixAndPrune(secret, secret/32)
		secret = mixAndPrune(secret, secret*2048)
		prices = append(prices, secret%10)
	}
	return secret, prices
}

func sell(prices, instructions []int) int {
	difs := []int{}
main:
	for i := 1; i < len(prices); i++ {
		difs = append(difs, prices[i]-prices[i-1])
		if len(difs) == 4 {
			for j := range difs {
				if difs[j] != instructions[j] {
					difs = difs[1:]
					continue main
				}
			}
			return prices[i]
		}
	}
	return 0
}

func genInstructionCache(prices []int) [][][][]int {
	cache := make([][][][]int, 19)
	for a := range cache {
		cache[a] = make([][][]int, 19)
		for b := range cache[a] {
			cache[a][b] = make([][]int, 19)
			for c := range cache[a][b] {
				cache[a][b][c] = make([]int, 19)
				for d := range cache[a][b][c] {
					cache[a][b][c][d] = -1
				}
			}
		}
	}
	for i := 4; i < len(prices); i++ {
		a := prices[i-3] - prices[i-4] + 9
		b := prices[i-2] - prices[i-3] + 9
		c := prices[i-1] - prices[i-2] + 9
		d := prices[i] - prices[i-1] + 9
		if cache[a][b][c][d] == -1 {
			cache[a][b][c][d] = prices[i]
		}
	}
	return cache
}
