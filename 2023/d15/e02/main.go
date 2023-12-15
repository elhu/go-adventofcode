package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Lens struct {
	fLen  int
	label string
}

type Box struct {
	lenses []Lens
}

type Step struct {
	label string
	op    byte
	param int
}

func hash(inst string) int {
	h := 0
	for _, c := range inst {
		h += int(c)
		h *= 17
		h %= 256
	}
	return h
}

var exp = regexp.MustCompile(`([a-z]+)([=\-])([0-9]?)`)

func parseStep(step string) Step {
	var s Step
	matches := exp.FindStringSubmatch(step)
	s.label = matches[1]
	s.op = matches[2][0]
	if s.op == '=' {
		s.param = int(matches[3][0] - '0')
	}
	return s
}

func index(lenses []Lens, label string) int {
	for i, l := range lenses {
		if l.label == label {
			return i
		}
	}
	return -1
}

func remove(step Step, boxes *[256]Box) {
	boxIndex := hash(step.label)
	box := boxes[boxIndex]
	index := index(box.lenses, step.label)
	if index != -1 {
		box.lenses = append(box.lenses[:index], box.lenses[index+1:]...)
	}
	boxes[boxIndex] = box
}

func add(step Step, boxes *[256]Box) {
	boxIndex := hash(step.label)
	box := boxes[boxIndex]
	index := index(box.lenses, step.label)
	if index == -1 {
		box.lenses = append(box.lenses, Lens{step.param, step.label})
	} else {
		box.lenses[index].fLen = step.param
	}
	boxes[boxIndex] = box
}

func computeScore(boxes [256]Box) int {
	score := 0
	for i, box := range boxes {
		for j, lens := range box.lenses {
			score += (i + 1) * (j + 1) * lens.fLen
		}
	}
	return score
}

func solve(instructions []string) int {
	var boxes [256]Box
	for _, inst := range instructions {
		step := parseStep(inst)
		if step.op == '-' {
			remove(step, &boxes)
		} else if step.op == '=' {
			add(step, &boxes)
		} else {
			panic("WTF")
		}
	}
	return computeScore(boxes)
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	instructions := strings.Split(data, ",")
	fmt.Println(solve(instructions))
}
