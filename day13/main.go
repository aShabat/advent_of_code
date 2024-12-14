package main

import (
	"advent/util"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	lines := util.ReadFile("day13/input.txt")
	games := []claw{}
	for len(lines) > 0 {
		games = append(games, parseClaw(lines))
		lines = lines[3:]
		if len(lines) > 0 {
			lines = lines[1:]
		}
	}

	out1 := 0
	out2 := 0
	for _, game := range games {
		out1 += game.solve1()
		out2 += game.solve2()
	}
	fmt.Println(out1, out2)
}

type vector struct {
	x int
	y int
}
type claw struct {
	a     vector
	b     vector
	prize vector
}

func parseClaw(lines []string) claw {
	reA := regexp.MustCompile(`Button A\: X\+(\d+), Y\+(\d+)`)
	submatchA := reA.FindAllStringSubmatch(lines[0], -1)
	xA, _ := strconv.Atoi(submatchA[0][1])
	yA, _ := strconv.Atoi(submatchA[0][2])
	A := vector{x: xA, y: yA}

	reB := regexp.MustCompile(`Button B\: X\+(\d+), Y\+(\d+)`)
	submatchB := reB.FindAllStringSubmatch(lines[1], -1)
	xB, _ := strconv.Atoi(submatchB[0][1])
	yB, _ := strconv.Atoi(submatchB[0][2])
	B := vector{x: xB, y: yB}

	reP := regexp.MustCompile(`Prize\: X\=(\d+), Y\=(\d+)`)
	submatchP := reP.FindAllStringSubmatch(lines[2], -1)
	xP, _ := strconv.Atoi(submatchP[0][1])
	yP, _ := strconv.Atoi(submatchP[0][2])
	Prize := vector{x: xP, y: yP}

	return claw{a: A, b: B, prize: Prize}
}

func (this claw) solve1() int {
	tokens := 0
	for a := range 101 {
		for b := range 101 {
			if a*this.a.x+b*this.b.x == this.prize.x &&
				a*this.a.y+b*this.b.y == this.prize.y &&
				(3*a+b < tokens || tokens == 0) {
				tokens = 3*a + b
			}
		}
	}
	return tokens
}

func (this claw) solve2() int {
	this.prize.x += 1e13
	this.prize.y += 1e13

	aN := this.prize.x*this.b.y - this.prize.y*this.b.x
	aD := this.a.x*this.b.y - this.a.y*this.b.x
	if aN%aD != 0 {
		return 0
	}

	bN := this.prize.x*this.a.y - this.prize.y*this.a.x
	bD := this.b.x*this.a.y - this.b.y*this.a.x
	if bN%bD != 0 {
		return 0
	}

	return 3*(aN/aD) + (bN / bD)
}
