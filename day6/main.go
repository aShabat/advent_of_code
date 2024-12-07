package main

import (
	"advent/util"
	"fmt"
)

func main() {
	lines := util.ReadFile("day6/input.txt")
	Map := make([][]int, len(lines))
	directions := [][]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
	x, y := -1, -1
	dir := -1
	for i, line := range lines {
		Map[i] = make([]int, len(line))
		for j, l := range line {
			if l == '#' {
				Map[i][j] = -1
			} else {
				Map[i][j] = 0
				if l != '.' {
					x, y = i, j
					switch l {
					case '^':
						dir = 0
					case '>':
						dir = 1
					case 'v':
						dir = 2
					case '<':
						dir = 3
					default:
						panic(0)
					}
				}
			}
		}
	}

	out2 := 0
	for i := range Map {
	testObstacle:
		for j := range Map[i] {
			if Map[i][j] == -1 || i == x && j == y {
				continue
			}

			X, Y, Dir := x, y, dir
			dirMap := make([][]map[int]struct{}, len(Map))
			for X >= 0 && X < len(Map) && Y >= 0 && Y < len(Map[X]) {
				if len(dirMap[X]) == 0 {
					dirMap[X] = make([]map[int]struct{}, len(Map[X]))
				}
				if dirMap[X][Y] == nil {
					dirMap[X][Y] = map[int]struct{}{}
				}
				if _, ok := dirMap[X][Y][Dir]; ok {
					out2++
					fmt.Println(i, j, X, Y, Dir)
					continue testObstacle
				} else {
					dirMap[X][Y][Dir] = struct{}{}
				}

				newX, newY := X+directions[Dir][0], Y+directions[Dir][1]
				if newX < 0 || newX >= len(Map) || newY < 0 || newY >= len(Map[X]) {
					continue testObstacle
				}
				if Map[newX][newY] == -1 || newX == i && newY == j {
					Dir = (Dir + 1) % 4
				} else {
					X, Y = newX, newY
				}
			}
		}
	}

	out1 := 0
	for x >= 0 && x < len(Map) && y >= 0 && y < len(Map[x]) {
		if Map[x][y] == 0 {
			Map[x][y] = 1
			out1++
		}

		newX, newY := x+directions[dir][0], y+directions[dir][1]
		if newX < 0 || newX >= len(Map) || newY < 0 || newY >= len(Map[newX]) {
			break
		}
		if Map[newX][newY] == -1 {
			dir = (dir + 1) % 4
		} else {
			x, y = newX, newY
		}
	}

	fmt.Println(out1, out2)
}
