package main

import (
	"advent/util"
	"fmt"
	"maps"
	"sort"
	"strings"
)

func main() {
	lines := util.ReadFile("day23/input.txt")
	conns := map[string]map[string]struct{}{}
	comps := map[string]struct{}{}
	for _, line := range lines {
		lineSplit := strings.Split(line, "-")
		first, second := lineSplit[0], lineSplit[1]
		conns = addConnection(conns, first, second)
		comps[first] = struct{}{}
		comps[second] = struct{}{}
	}
	triplets := map[[3]string]struct{}{}
	out1 := 0
	for a := range conns {
		for b := range conns[a] {
		tripleSearch:
			for c := range conns[b] {
				if a == c {
					continue
				}
				if _, ok := conns[a][c]; !ok {
					continue
				}
				for _, triplet := range [][3]string{
					{a, b, c},
					{a, c, b},
					{b, a, c},
					{b, c, a},
					{c, a, b},
					{c, b, a},
				} {
					if _, ok := triplets[triplet]; ok {
						continue tripleSearch
					}
				}
				if tStart(a) || tStart(b) || tStart(c) {
					triplets[[3]string{a, b, c}] = struct{}{}
					out1++
				}
			}
		}
	}
	fmt.Println(out1)
	sets := []map[string]struct{}{{}}
	for {
		newSets := []map[string]struct{}{}
		for comp := range conns {
		search:
			for _, set := range sets {
				for sComp := range set {
					if _, ok := conns[comp][sComp]; !ok || comp <= sComp {
						continue search
					}
				}
				newSet := maps.Clone(set)
				newSet[comp] = struct{}{}
				newSets = append(newSets, newSet)
			}
		}
		if len(newSets) == 0 {
			break
		}
		sets = newSets
	}
	set := []string{}
	for comp := range sets[0] {
		set = append(set, comp)
	}
	sort.Strings(set)
	out2 := ""
	for _, comp := range set {
		out2 += "," + comp
	}
	fmt.Println(out2[1:])
}

func addConnection(conns map[string]map[string]struct{}, first, second string) map[string]map[string]struct{} {
	if _, ok := conns[first]; !ok {
		conns[first] = map[string]struct{}{}
	}
	if _, ok := conns[first][second]; !ok {
		conns[first][second] = struct{}{}
	}
	if _, ok := conns[second]; !ok {
		conns[second] = map[string]struct{}{}
	}
	if _, ok := conns[second][first]; !ok {
		conns[second][first] = struct{}{}
	}

	return conns
}

func tStart(comp string) bool {
	return strings.HasPrefix(comp, "t")
}
