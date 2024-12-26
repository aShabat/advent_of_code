package main

import (
	"advent/util"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	lines := util.ReadFile("day17/input.txt")
	A, _ := strconv.Atoi(strings.Split(lines[0], " ")[2])
	B, _ := strconv.Atoi(strings.Split(lines[1], " ")[2])
	C, _ := strconv.Atoi(strings.Split(lines[2], " ")[2])
	comp := &computer{A: A, B: B, C: C, pointer: 0, output: []int{}}
	instructions := util.Split(strings.Split(lines[4], " ")[1], ",")
	comp.run(instructions)
	for _, num := range comp.output {
		fmt.Print(num, ",")
	}
	fmt.Println()

	out2 := 0
	for i := range len(instructions) {
		out2 *= 8
		for ; ; out2++ {
			comp.A = out2
			comp.B = 0
			comp.C = 0
			comp.output = []int{}
			comp.pointer = 0
			comp.run(instructions)
			if slices.Equal(comp.output[len(comp.output)-i-1:], instructions[len(instructions)-i-1:]) {
				break
			}
		}
	}
	fmt.Println(out2)
}

type computer struct {
	A, B, C int
	pointer int
	output  []int
}

func (this computer) combo(operand int) int {
	switch operand {
	case 4:
		return this.A
	case 5:
		return this.B
	case 6:
		return this.C
	default:
		return operand
	}
}

func (this *computer) instruction(operation, operand int) {
	switch operation {
	case 0:
		this.A = this.A / power(this.combo(operand))
	case 1:
		this.B = this.B ^ operand
	case 2:
		this.B = this.combo(operand) % 8
	case 3:
		if this.A != 0 {
			this.pointer = operand - 2
		}
	case 4:
		this.B = this.B ^ this.C
	case 5:
		this.output = append(this.output, this.combo(operand)%8)
	case 6:
		this.B = this.A / power(this.combo(operand))
	case 7:
		this.C = this.A / power(this.combo(operand))
	}
}

func (this *computer) run(instructions []int) {
	for this.pointer < len(instructions) {
		this.instruction(instructions[this.pointer], instructions[this.pointer+1])
		this.pointer += 2
	}
}

func power(p int) int {
	if p == 0 {
		return 1
	}
	return 2 * power(p-1)
}
