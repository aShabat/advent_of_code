package main

import (
	"advent/util"
	"fmt"
	"slices"
)

func main() {
	lines := util.ReadFile("day15/input.txt")
	warehouse := [][]int{}
	var rRow, rCol int
	for {
		line := lines[0]
		lines = lines[1:]
		if line == "" {
			break
		}
		warehouse = append(warehouse, make([]int, len(line)))
		for i, char := range line {
			if char == '#' {
				warehouse[len(warehouse)-1][i] = -1
			} else if char == '.' {
				warehouse[len(warehouse)-1][i] = 0
			} else if char == 'O' {
				warehouse[len(warehouse)-1][i] = 1
			} else {
				warehouse[len(warehouse)-1][i] = 0
				rRow, rCol = len(warehouse)-1, i
			}
		}
	}
	for _, line := range lines {
		for _, command := range line {
			var direction []int
			switch command {
			case '<':
				direction = []int{0, -1}
			case '>':
				direction = []int{0, 1}
			case '^':
				direction = []int{-1, 0}
			case 'v':
				direction = []int{1, 0}
			}

			lookAhead := 1
			for warehouse[rRow+direction[0]*lookAhead][rCol+direction[1]*lookAhead] == 1 {
				lookAhead++
			}
			if warehouse[rRow+direction[0]*lookAhead][rCol+direction[1]*lookAhead] == 0 {
				if lookAhead != 1 {
					warehouse[rRow+direction[0]][rCol+direction[1]] = 0
					warehouse[rRow+direction[0]*lookAhead][rCol+direction[1]*lookAhead] = 1
				}
				rRow += direction[0]
				rCol += direction[1]
			}
		}
	}

	out1 := 0
	for row := range warehouse {
		for col := range warehouse[row] {
			if warehouse[row][col] == 1 {
				out1 += col + 100*row
			}
		}
	}
	fmt.Println(out1)

	lines = util.ReadFile("day15/input.txt")
	warehouse = [][]int{}
	for {
		line := lines[0]
		lines = lines[1:]
		if line == "" {
			break
		}
		warehouse = append(warehouse, make([]int, 2*len(line)))
		for i, char := range line {
			if char == '#' {
				warehouse[len(warehouse)-1][2*i] = -1
				warehouse[len(warehouse)-1][2*i+1] = -1
			} else if char == '.' {
				warehouse[len(warehouse)-1][2*i] = 0
				warehouse[len(warehouse)-1][2*i+1] = 0
			} else if char == 'O' {
				warehouse[len(warehouse)-1][2*i] = 1
				warehouse[len(warehouse)-1][2*i+1] = 2
			} else {
				warehouse[len(warehouse)-1][2*i] = 0
				warehouse[len(warehouse)-1][2*i+1] = 0
				rRow, rCol = len(warehouse)-1, 2*i
			}
		}
	}
	for _, line := range lines {
	robotFor:
		for _, command := range line {
			switch command {
			case '<':
				lookAhead := 1
				for warehouse[rRow][rCol-lookAhead] >= 1 {
					lookAhead += 2
				}
				if warehouse[rRow][rCol-lookAhead] == 0 {
					if lookAhead != 1 {
						for i := range lookAhead {
							warehouse[rRow][rCol-lookAhead+i] = 1 + (i % 2)
						}
						warehouse[rRow][rCol-1] = 0
					}
					rCol--
				}
			case '>':
				lookAhead := 1
				for warehouse[rRow][rCol+lookAhead] >= 1 {
					lookAhead += 2
				}
				if warehouse[rRow][rCol+lookAhead] == 0 {
					if lookAhead != 1 {
						for i := range lookAhead {
							warehouse[rRow][rCol+lookAhead-i] = 2 - (i % 2)
						}
						warehouse[rRow][rCol+1] = 0
					}
					rCol++
				}
			case '^':
				var lookAhead [][]int
				if warehouse[rRow-1][rCol] >= 1 {
					lookAhead = [][]int{{rRow - 1, rCol}}
				} else if warehouse[rRow-1][rCol] == -1 {
					continue robotFor
				}
				for i := 0; i < len(lookAhead); i++ {
					if warehouse[lookAhead[i][0]][lookAhead[i][1]] == 2 {
						lookAhead[i][1]--
					}
					if warehouse[lookAhead[i][0]-1][lookAhead[i][1]] == -1 || warehouse[lookAhead[i][0]-1][lookAhead[i][1]+1] == -1 {
						continue robotFor
					}
					next := warehouse[lookAhead[i][0]-1][lookAhead[i][1]]
					if next == 1 {
						lookAhead = append(lookAhead, []int{lookAhead[i][0] - 1, lookAhead[i][1]})
					} else if next == 2 {
						lookAhead = append(lookAhead, []int{lookAhead[i][0] - 1, lookAhead[i][1] - 1})
					}
					next = warehouse[lookAhead[i][0]-1][lookAhead[i][1]+1]
					if next == 1 {
						lookAhead = append(lookAhead, []int{lookAhead[i][0] - 1, lookAhead[i][1] + 1})
					}
				}
				slices.Reverse(lookAhead)
				for _, box := range lookAhead {
					warehouse[box[0]][box[1]] = 0
					warehouse[box[0]][box[1]+1] = 0
					warehouse[box[0]-1][box[1]] = 1
					warehouse[box[0]-1][box[1]+1] = 2
				}
				rRow--
			case 'v':
				var lookAhead [][]int
				if warehouse[rRow+1][rCol] >= 1 {
					lookAhead = [][]int{{rRow + 1, rCol}}
				} else if warehouse[rRow+1][rCol] == -1 {
					continue robotFor
				}
				for i := 0; i < len(lookAhead); i++ {
					if warehouse[lookAhead[i][0]][lookAhead[i][1]] == 2 {
						lookAhead[i][1]--
					}
					if warehouse[lookAhead[i][0]+1][lookAhead[i][1]] == -1 || warehouse[lookAhead[i][0]+1][lookAhead[i][1]+1] == -1 {
						continue robotFor
					}
					next := warehouse[lookAhead[i][0]+1][lookAhead[i][1]]
					if next == 1 {
						lookAhead = append(lookAhead, []int{lookAhead[i][0] + 1, lookAhead[i][1]})
					} else if next == 2 {
						lookAhead = append(lookAhead, []int{lookAhead[i][0] + 1, lookAhead[i][1] - 1})
					}
					next = warehouse[lookAhead[i][0]+1][lookAhead[i][1]+1]
					if next == 1 {
						lookAhead = append(lookAhead, []int{lookAhead[i][0] + 1, lookAhead[i][1] + 1})
					}
				}
				slices.Reverse(lookAhead)
				for _, box := range lookAhead {
					warehouse[box[0]][box[1]] = 0
					warehouse[box[0]][box[1]+1] = 0
					warehouse[box[0]+1][box[1]] = 1
					warehouse[box[0]+1][box[1]+1] = 2
				}
				rRow++
			}
		}
	}

	out2 := 0
	for row := range warehouse {
		for col := range warehouse[row] {
			if warehouse[row][col] == 1 {
				out2 += col + 100*row
			}
		}
	}
	fmt.Println(out2)
}
