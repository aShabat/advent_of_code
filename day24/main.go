package main

import (
	"advent/util"
	"fmt"
	"regexp"
	"slices"
	"strconv"
)

func main() {
	lines := util.ReadFile("day24/input.txt")
	wires := map[string]int{}
	readWireRX := regexp.MustCompile(`(.+)\: (.)`)
	for {
		line := lines[0]
		lines = lines[1:]
		if line == "" {
			break
		}
		parse := readWireRX.FindAllStringSubmatch(line, -1)[0]
		wires[parse[1]], _ = strconv.Atoi(parse[2])
	}
	gates := [][]string{}
	readGateRX := regexp.MustCompile(`(.+) (.+) (.+) -> (.+)`)
	requiredInGates := map[string][]int{}
	for _, line := range lines {
		parse := readGateRX.FindAllStringSubmatch(line, -1)[0][1:]
		gates = append(gates, parse)
		if _, ok := wires[parse[0]]; !ok {
			wires[parse[0]] = -1
		}
		if _, ok := requiredInGates[parse[0]]; !ok {
			requiredInGates[parse[0]] = []int{}
		}
		requiredInGates[parse[0]] = append(requiredInGates[parse[0]], len(gates)-1)
		if _, ok := wires[parse[2]]; !ok {
			wires[parse[2]] = -1
		}
		if _, ok := requiredInGates[parse[2]]; !ok {
			requiredInGates[parse[2]] = []int{}
		}
		requiredInGates[parse[2]] = append(requiredInGates[parse[2]], len(gates)-1)
		if _, ok := wires[parse[3]]; !ok {
			wires[parse[3]] = -1
		}
	}
	gateCheckStack := []int{}
	for key, gates := range requiredInGates {
		if wires[key] != -1 {
			gateCheckStack = append(gateCheckStack, gates...)
		}
	}
	for len(gateCheckStack) > 0 {
		now := gateCheckStack[0]
		gateCheckStack = gateCheckStack[1:]
		left, right, gate, result := gates[now][0], gates[now][2], gates[now][1], gates[now][3]
		if wires[result] == -1 && wires[left] != -1 && wires[right] != -1 {
			switch gate {
			case "AND":
				wires[result] = wires[left] & wires[right]
			case "OR":
				wires[result] = wires[left] | wires[right]
			case "XOR":
				wires[result] = wires[left] ^ wires[right]
			}
		}
		gateCheckStack = append(gateCheckStack, requiredInGates[result]...)
	}
	outBytes := []int{}
	for i := 0; ; i++ {
		key := "z" + fmt.Sprintf("%02d", i)
		if byte, ok := wires[key]; ok {
			outBytes = append(outBytes, byte)
		} else {
			break
		}
	}
	out1 := 0
	slices.Reverse(outBytes)
	for _, byte := range outBytes {
		out1 = 2*out1 + byte
	}
	fmt.Println(out1)
}
