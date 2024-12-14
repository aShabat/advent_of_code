package main

import (
	"advent/util"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	lines := util.ReadFile("day14/input.txt")
	robots := make([][]int, len(lines))
	re := regexp.MustCompile(`p\=(.+),(.+) v\=(.+),(.+)`)
	for i, line := range lines {
		numStrings := re.FindAllStringSubmatch(line, -1)
		for j := range 4 {
			num, _ := strconv.Atoi(numStrings[0][j+1])
			robots[i] = append(robots[i], num)
		}
	}
	mapWidth, mapHeight := 101, 103
	middleWidth, middleHeight := (mapWidth-1)/2, (mapHeight-1)/2
	out1 := 0
	quadrants := []int{0, 0, 0, 0}
	for _, robot := range robots {
		x := (robot[0] + 100*(mapWidth+robot[2])) % mapWidth
		y := (robot[1] + 100*(mapHeight+robot[3])) % mapHeight
		if x > middleWidth && y > middleHeight {
			quadrants[0]++
		} else if x > middleWidth && y < middleHeight {
			quadrants[1]++
		} else if x < middleWidth && y < middleHeight {
			quadrants[2]++
		} else if x < middleWidth && y > middleHeight {
			quadrants[3]++
		}
	}
	out1 = quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]

	fmt.Println(out1)
	minSpace := mapWidth * mapHeight
	for i := 0; ; i++ {
		robotMap := make([][]bool, mapHeight)
		for i := range robotMap {
			robotMap[i] = make([]bool, mapWidth)
		}
		for _, robot := range robots {
			x := (robot[0] + i*(mapWidth+robot[2])) % mapWidth
			y := (robot[1] + i*(mapHeight+robot[3])) % mapHeight
			robotMap[y][x] = true
		}
		dx, countx := -1, 0
		for countx < 300 {
			countx = 0
			dx++
			for i := middleWidth - dx; i <= middleWidth+dx; i++ {
				for j := range mapHeight {
					if robotMap[j][i] {
						countx++
					}
				}
			}
		}
		dy, county := -1, 0
		for county < 300 {
			county = 0
			dy++
			for j := middleHeight - dy; j <= middleHeight+dy; j++ {
				for i := range mapWidth {
					if robotMap[j][i] {
						county++
					}
				}
			}
		}
		if dx*dy < minSpace {
			minSpace = dx * dy
			fmt.Println(i, dx, dy)
			for _, line := range robotMap {
				for _, r := range line {
					if r {
						fmt.Print("*")
					} else {
						fmt.Print(".")
					}
				}
				fmt.Print("\n")
			}
			fmt.Scanln()
		}
	}
}
