package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
)

type sortBytes []byte

func (s sortBytes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortBytes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortBytes) Len() int {
	return len(s)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type step struct {
	before map[byte]struct{}
	name   byte
}

var exp = regexp.MustCompile(`Step (\w) must be finished before step (\w) can begin`)

func fetchStep(steps map[byte]*step, name byte) *step {
	if step, exists := steps[name]; exists {
		return step
	}
	step := &step{make(map[byte]struct{}), name}
	steps[name] = step
	return step
}

func findReadySteps(steps map[byte]*step) []byte {
	res := make([]byte, 0)
	for name, step := range steps {
		if len(step.before) == 0 {
			res = append(res, name)
		}
	}
	return res
}

func solve(steps map[byte]*step) []byte {
	res := make([]byte, 0)
	for ready := findReadySteps(steps); len(ready) > 0; ready = findReadySteps(steps) {
		sort.Sort(sortBytes(ready))
		s := ready[0]
		res = append(res, s)
		delete(steps, s)
		for _, step := range steps {
			delete(step.before, s)
		}
	}
	return res
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	data = bytes.TrimSuffix(data, []byte{'\n'})
	steps := make(map[byte]*step)
	lines := bytes.Split(data, []byte{'\n'})
	for _, l := range lines {
		match := exp.FindStringSubmatch(string(l))
		before := []byte(match[1])[0]
		after := []byte(match[2])[0]
		afterStep := fetchStep(steps, after)
		fetchStep(steps, before)
		afterStep.before[before] = struct{}{}
	}
	fmt.Println(string(solve(steps)))
}
