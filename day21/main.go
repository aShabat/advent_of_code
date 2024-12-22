package main

import (
	"advent/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := util.ReadFile("day21/input.txt")
	out1 := 0
	out2 := 0
	for _, line := range lines {
		line = "A" + line
		instructions := []string{""}
		for i := range len(line) - 1 {
			insNext := []string{}
			for _, option := range numToNum(line[i], line[i+1]) {
				for _, i := range instructions {
					insNext = append(insNext, i+option+"A")
				}
			}
			instructions = insNext
		}
		meta1 := metaLen(instructions[0], 2)
		meta2 := metaLen(instructions[0], 25)
		for _, instruction := range instructions[1:] {
			if other := metaLen(instruction, 2); meta1 > other {
				meta1 = other
			}
			if other := metaLen(instruction, 25); meta2 > other {
				meta2 = other
			}
		}
		num, _ := strconv.Atoi(strings.Split(line, "A")[1])
		out1 += meta1 * num
		out2 += meta2 * num
	}
	fmt.Println(out1, out2)
}

type metaKey struct {
	instruction string
	depth       int
}

var metaCache = map[metaKey]int{}

func metaLen(instruction string, depth int) int {
	if out, ok := metaCache[metaKey{instruction, depth}]; ok {
		return out
	}
	var out int
	if depth == 0 {
		out = len(instruction)
	} else if strings.Index(instruction, "A") != len(instruction)-1 {
		for _, s := range strings.SplitAfter(instruction, "A") {
			out += metaLen(s, depth)
		}
	} else {
		instruction = "A" + instruction
		insUp := []string{""}
		for i := range len(instruction) - 1 {
			insUpNext := []string{}
			for _, option := range keyToKey(instruction[i], instruction[i+1]) {
				for _, iu := range insUp {
					insUpNext = append(insUpNext, iu+option+"A")
				}
			}
			insUp = insUpNext
		}
		out = metaLen(insUp[0], depth-1)
		for _, iu := range insUp[1:] {
			if other := metaLen(iu, depth-1); other < out {
				out = other
			}
		}
		instruction = instruction[1:]
	}

	metaCache[metaKey{instruction, depth}] = out
	return out
}

var numToNumCache map[[2]byte][]string = map[[2]byte][]string{}

func numToNum(start, finish byte) []string {
	if out, ok := numToNumCache[[2]byte{start, finish}]; ok {
		return out
	}
	positions := map[byte][]int{
		'0': {0, 1},
		'1': {1, 0},
		'2': {1, 1},
		'3': {1, 2},
		'4': {2, 0},
		'5': {2, 1},
		'6': {2, 2},
		'7': {3, 0},
		'8': {3, 1},
		'9': {3, 2},
		'A': {0, 2},
	}
	sPos, fPos := positions[start], positions[finish]
	out := []string{}
	if sPos[0] == fPos[0] && sPos[1] == fPos[1] {
		out = append(out, "")
	} else if sPos[0] == fPos[0] {
		if sPos[1] > fPos[1] {
			out = append(out, helper("<", sPos[1]-fPos[1]))
		} else {
			out = append(out, helper(">", fPos[1]-sPos[1]))
		}
	} else if sPos[1] == fPos[1] {
		if sPos[0] > fPos[0] {
			out = append(out, helper("v", sPos[0]-fPos[0]))
		} else {
			out = append(out, helper("^", fPos[0]-sPos[0]))
		}
	} else {
		if sPos[0] > fPos[0] {
			if sPos[1] > fPos[1] {
				out = append(out, helper("v", sPos[0]-fPos[0])+helper("<", sPos[1]-fPos[1]))
				out = append(out, helper("<", sPos[1]-fPos[1])+helper("v", sPos[0]-fPos[0]))
			} else {
				out = append(out, helper(">", fPos[1]-sPos[1])+helper("v", sPos[0]-fPos[0]))
				if sPos[1] != 0 || fPos[0] != 0 {
					out = append(out, helper("v", sPos[0]-fPos[0])+helper(">", fPos[1]-sPos[1]))
				}
			}
		} else {
			if sPos[1] < fPos[1] {
				out = append(out, helper("^", fPos[0]-sPos[0])+helper(">", fPos[1]-sPos[1]))
				out = append(out, helper(">", fPos[1]-sPos[1])+helper("^", fPos[0]-sPos[0]))
			} else {
				out = append(out, helper("^", fPos[0]-sPos[0])+helper("<", sPos[1]-fPos[1]))
				if sPos[0] != 0 || fPos[1] != 0 {
					out = append(out, helper("<", sPos[1]-fPos[1])+helper("^", fPos[0]-sPos[0]))
				}
			}
		}
	}

	numToNumCache[[2]byte{start, finish}] = out
	return out
}

var keyToKeyCache map[[2]byte][]string = map[[2]byte][]string{}

func keyToKey(start, finish byte) []string {
	if out, ok := keyToKeyCache[[2]byte{start, finish}]; ok {
		return out
	}
	out := []string{}
	positions := map[byte][]int{
		'A': {0, 2},
		'^': {0, 1},
		'<': {1, 0},
		'v': {1, 1},
		'>': {1, 2},
	}
	sPos, fPos := positions[start], positions[finish]
	if sPos[0] == fPos[0] && sPos[1] == fPos[1] {
		out = append(out, "")
	} else if sPos[0] == fPos[0] {
		if sPos[1] > fPos[1] {
			out = append(out, helper("<", sPos[1]-fPos[1]))
		} else {
			out = append(out, helper(">", fPos[1]-sPos[1]))
		}
	} else if sPos[1] == fPos[1] {
		if sPos[0] > fPos[0] {
			out = append(out, helper("^", sPos[0]-fPos[0]))
		} else {
			out = append(out, helper("v", fPos[0]-sPos[0]))
		}
	} else {
		if sPos[0] > fPos[0] {
			if sPos[1] > fPos[1] {
				out = append(out, helper("^", sPos[0]-fPos[0])+helper("<", sPos[1]-fPos[1]))
				out = append(out, helper("<", sPos[1]-fPos[1])+helper("^", sPos[0]-fPos[0]))
			} else {
				out = append(out, helper(">", fPos[1]-sPos[1])+helper("^", sPos[0]-fPos[0]))
				if sPos[1] != 0 || fPos[0] != 0 {
					out = append(out, helper("^", sPos[0]-fPos[0])+helper(">", fPos[1]-sPos[1]))
				}
			}
		} else {
			if sPos[1] < fPos[1] {
				out = append(out, helper("v", fPos[0]-sPos[0])+helper(">", fPos[1]-sPos[1]))
				out = append(out, helper(">", fPos[1]-sPos[1])+helper("v", fPos[0]-sPos[0]))
			} else {
				out = append(out, helper("v", fPos[0]-sPos[0])+helper("<", sPos[1]-fPos[1]))
				if sPos[0] != 0 || fPos[1] != 0 {
					out = append(out, helper("<", sPos[1]-fPos[1])+helper("v", fPos[0]-sPos[0]))
				}
			}
		}
	}

	keyToKeyCache[[2]byte{start, finish}] = out
	return out
}

func helper(char string, count int) string {
	out := ""
	for range count {
		out += char
	}
	return out
}
