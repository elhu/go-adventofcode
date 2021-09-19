// No code change, changed the input instead

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func atoi(str string) int {
	i, e := strconv.Atoi(str)
	check(e)
	return i
}

type Input struct {
	gateType string
	argLeft  interface{}
	argRight interface{}
}

type Wire struct {
	val   int
	name  string
	input Input
}

func isNumber(str string) bool {
	_, e := strconv.Atoi(str)
	return e == nil
}

var parseExp = regexp.MustCompile(`([a-z0-9]+ )?([A-Z]+ )?([a-z0-9]+) -> ([a-z0-9]+)`)

func computeInput(wires map[string]*Wire, name string) int {
	// fmt.Printf("Processing input for %s\n", name)
	if wires[name].val != -1 {
		return wires[name].val
	}
	input := wires[name].input
	val := -1
	switch input.gateType {
	case "VAL":
		r, ok := input.argRight.(int)
		if ok {
			val = r
		} else {
			val = computeInput(wires, input.argRight.(*Wire).name)
		}
	case "NOT":
		r, ok := input.argRight.(int)
		if ok {
			val = not(r)
		} else {
			val = not(computeInput(wires, input.argRight.(*Wire).name))
		}
	case "AND":
		l, ok := input.argLeft.(int)
		if !ok {
			l = computeInput(wires, input.argLeft.(*Wire).name)
		}
		r, ok := input.argRight.(int)
		if !ok {
			r = computeInput(wires, input.argRight.(*Wire).name)
		}
		val = r & l
	case "OR":
		l, ok := input.argLeft.(int)
		if !ok {
			l = computeInput(wires, input.argLeft.(*Wire).name)
		}
		r, ok := input.argRight.(int)
		if !ok {
			r = computeInput(wires, input.argRight.(*Wire).name)
		}
		val = r | l
	case "LSHIFT":
		l, ok := input.argLeft.(int)
		if !ok {
			l = computeInput(wires, input.argLeft.(*Wire).name)
		}
		r, ok := input.argRight.(int)
		if !ok {
			r = computeInput(wires, input.argRight.(*Wire).name)
		}
		val = l << r
	case "RSHIFT":
		l, ok := input.argLeft.(int)
		if !ok {
			l = computeInput(wires, input.argLeft.(*Wire).name)
		}
		r, ok := input.argRight.(int)
		if !ok {
			r = computeInput(wires, input.argRight.(*Wire).name)
		}
		val = l >> r
	}
	if val != -1 {
		wires[name].val = val
		return val
	}
	fmt.Println(input)
	panic("WTF")
}

func solve(wires map[string]*Wire, target string) int {
	wire := wires[target]
	wire.val = computeInput(wires, target)
	return wire.val
}

func parseInput(input []string) map[string]*Wire {
	wires := make(map[string]*Wire)
	wires["NULL"] = &Wire{}
	for _, l := range input {
		matches := parseExp.FindStringSubmatch(l)
		for i, v := range matches {
			matches[i] = strings.TrimRight(v, " ")
		}
		valLeft, kind, valRight, name := matches[1], matches[2], matches[3], matches[4]
		if kind == "" {
			kind = "VAL"
		}
		if _, found := wires[name]; !found {
			wires[name] = &Wire{name: name, val: -1}
		}
		var argLeft interface{}
		argLeft = wires["NULL"]
		if valLeft != "" {
			if isNumber(valLeft) {
				argLeft = atoi(valLeft)
			} else {
				if _, found := wires[valLeft]; !found {
					wires[valLeft] = &Wire{name: valLeft, val: -1}
				}
				argLeft = wires[valLeft]
			}
		}
		var argRight interface{}
		argRight = wires["NULL"]
		if valRight != "" {
			if isNumber(valRight) {
				argRight = atoi(valRight)
			} else {
				if _, found := wires[valRight]; !found {
					wires[valRight] = &Wire{name: valRight, val: -1}
				}
				argRight = wires[valRight]
			}
		}
		wires[name].input = Input{gateType: kind, argLeft: argLeft, argRight: argRight}
	}
	return wires
}

func not(val int) int {
	flipped := make([]byte, 16)
	valBin := strconv.FormatInt(int64(val), 2)
	valBin = strings.Repeat("0", len(flipped)-len(valBin)) + valBin
	for i, c := range valBin {
		if c == '1' {
			flipped[i] = '0'
		} else {
			flipped[i] = '1'
		}
	}
	res, err := strconv.ParseInt(string(flipped), 2, 32)
	check(err)
	return int(res)
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	wires := parseInput(input)
	fmt.Println(solve(wires, "a"))
}
