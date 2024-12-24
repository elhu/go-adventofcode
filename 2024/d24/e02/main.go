package main

import (
	"adventofcode/utils/sets"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

var OPS = map[string]func(int, int) int{
	"AND": func(a, b int) int {
		if a == 1 && b == 1 {
			return 1
		}
		return 0
	}, "OR": func(a, b int) int {
		if a == 1 || b == 1 {
			return 1
		}
		return 0
	}, "XOR": func(a, b int) int {
		if a != b {
			return 1
		}
		return 0
	},
}

type Gate struct {
	Op          func(int, int) int
	left, right string
	out         string
}

func parseWires(lines []string) map[string]int {
	wires := make(map[string]int)
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		wires[parts[0]] = atoi(parts[1])
	}
	return wires
}

func parseGates(lines []string) ([]*Gate, map[string]*sets.Set[*Gate]) {
	belongsTo := make(map[string]*sets.Set[*Gate])
	var gates []*Gate
	for _, line := range lines {
		parts := strings.Fields(line)
		g := &Gate{left: parts[0], right: parts[2], out: parts[4], Op: OPS[parts[1]]}
		if _, f := belongsTo[parts[0]]; !f {
			belongsTo[parts[0]] = sets.New[*Gate]()
		}
		if _, f := belongsTo[parts[2]]; !f {
			belongsTo[parts[2]] = sets.New[*Gate]()
		}
		belongsTo[parts[0]].Add(g)
		belongsTo[parts[2]].Add(g)
		gates = append(gates, g)
	}
	return gates, belongsTo
}

func findZWires(wires map[string]int, gates []*Gate) *sets.Set[string] {
	zWires := sets.New[string]()
	for k := range wires {
		if k[0] == 'z' {
			zWires.Add(k)
		}
	}
	for _, g := range gates {
		if g.out[0] == 'z' {
			zWires.Add(g.out)
		}
	}
	return zWires
}

func solve(wires map[string]int, gates []*Gate, belongsTo map[string]*sets.Set[*Gate]) int64 {
	zWires := findZWires(wires, gates)
	toProcess := sets.New[string]()
	for k := range wires {
		toProcess.Add(k)
	}
	for {
		if zWires.Len() == 0 {
			break
		}
		for _, p := range toProcess.Members() {
			if _, found := belongsTo[p]; !found {
				toProcess.Remove(p)
				continue
			}
			for _, g := range belongsTo[p].Members() {
				if _, found := wires[g.out]; found {
					continue
				}
				if _, found := wires[g.left]; !found {
					continue
				}
				if _, found := wires[g.right]; !found {
					continue
				}
				wires[g.out] = g.Op(wires[g.left], wires[g.right])
				belongsTo[g.left].Remove(g)
				toProcess.Add(g.out)
				if g.out[0] == 'z' {
					zWires.Remove(g.out)
				}
			}
			if belongsTo[p].Len() == 0 {
				toProcess.Remove(p)
			}
		}
		if toProcess.Len() == 0 {
			panic("NO SOLUTIONS")
		}
	}
	resW := findZWires(wires, gates).Members()
	sort.Sort(sort.Reverse(sort.StringSlice(resW)))
	res := make([]string, len(resW))
	for _, w := range resW {
		if wires[w] == 1 {
			res = append(res, "1")
		} else {
			res = append(res, "0")
		}
	}
	i, e := strconv.ParseInt(strings.Join(res, ""), 2, 64)
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	fmt.Println("Solved this by hand reordering the input file and figuring out where the order of operations wasn't correct. Good luck.")
	fmt.Println("See test.txt for results.\n")
	fmt.Println("cgr,hpc,hwk,qmd,tnt,z06,z31,z37")
}
