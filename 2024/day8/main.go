package main

import (
	"advent/util"
	"fmt"
)

func main() {
	lines := util.ReadFile("day8/input.txt")
	antennas := map[byte][][]int{}
	for y := range lines {
		for x := range lines[y] {
			b := lines[y][x]
			if b == '.' {
				continue
			}
			if _, ok := antennas[b]; !ok {
				antennas[b] = [][]int{}
			}
			antennas[b] = append(antennas[b], []int{x, y})
		}
	}

	isAntinode := make([][]bool, len(lines))
	for y := range isAntinode {
		isAntinode[y] = make([]bool, len(lines[y]))
	}
	out1 := 0
	for _, antennas := range antennas {
		for j := range antennas {
			for i := range j {
				x, y := 2*antennas[i][0]-antennas[j][0], 2*antennas[i][1]-antennas[j][1]
				if !(y < 0 || y >= len(isAntinode) || x < 0 || x >= len(isAntinode[y]) || isAntinode[y][x]) {
					isAntinode[y][x] = true
					out1++
				}
				x, y = 2*antennas[j][0]-antennas[i][0], 2*antennas[j][1]-antennas[i][1]
				if !(y < 0 || y >= len(isAntinode) || x < 0 || x >= len(isAntinode[y]) || isAntinode[y][x]) {
					isAntinode[y][x] = true
					out1++
				}
			}
		}
	}

	isAntinode = make([][]bool, len(lines))
	out2 := 0
	for y := range isAntinode {
		isAntinode[y] = make([]bool, len(lines[y]))
	}
	for _, antennas := range antennas {
		for j := range antennas {
			for i := range j {
				vector := []int{antennas[i][0] - antennas[j][0], antennas[i][1] - antennas[j][1]}
				vectorGid := gid(vector[0], vector[1])
				vector = []int{vector[0] / vectorGid, vector[1] / vectorGid}
				for n := 0; ; n++ {
					x, y := antennas[i][0]-n*vector[0], antennas[i][1]-n*vector[1]
					if y < 0 || y >= len(isAntinode) || x < 0 || x >= len(isAntinode[y]) {
						break
					} else if !isAntinode[y][x] {
						out2++
						isAntinode[y][x] = true
					}
				}
				for n := 1; ; n-- {
					x, y := antennas[i][0]-n*vector[0], antennas[i][1]-n*vector[1]
					if y < 0 || y >= len(isAntinode) || x < 0 || x >= len(isAntinode[y]) {
						break
					} else if !isAntinode[y][x] {
						out2++
						isAntinode[y][x] = true
					}
				}
			}
		}
	}

	fmt.Println(out1, out2)
}

func gid(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	if a < b {
		a, b = b, a
	}
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
