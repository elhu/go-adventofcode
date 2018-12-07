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
	before   map[byte]struct{}
	name     byte
	duration int
}

var exp = regexp.MustCompile(`Step (\w) must be finished before step (\w) can begin`)

func fetchStep(steps map[byte]*step, name byte) *step {
	if step, exists := steps[name]; exists {
		return step
	}
	step := &step{make(map[byte]struct{}), name, baseTime + int(name-'A') + 1}
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

const maxInt = int(^uint(0) >> 1)

func tickUntilFree(times []int) int {
	min := maxInt
	for _, t := range times {
		if t <= min && t > 0 {
			min = t
		}
	}
	for i := 0; i < len(times); i++ {
		times[i] -= min
	}
	return min
}

const baseTime = 60
const nbWorker = 5

func findFreeWorkers(workers []int) []int {
	res := make([]int, 0, 1)
	for i, w := range workers {
		if w <= 0 {
			res = append(res, i)
		}
	}
	return res
}

func solve(steps map[byte]*step) ([]byte, int) {
	res := make([]byte, 0)
	workers := make([]int, nbWorker)
	work := make([]byte, nbWorker)
	clockTime := 0
	for len(steps) > 0 {
		// Find steps that are ready
		ready := findReadySteps(steps)
		sort.Sort(sortBytes(ready))
		freeWorkers := findFreeWorkers(workers)
		// Assign steps to free workers
		for i := 0; i < len(ready) && i < len(freeWorkers); i++ {
			wID := freeWorkers[i]
			s := ready[i]
			workers[wID] = steps[s].duration
			work[wID] = s
			delete(steps, s)
		}
		// Tick until next job finishes
		timeShift := tickUntilFree(workers)
		clockTime += timeShift
		// Remove finished steps from steps pool
		for _, wID := range findFreeWorkers(workers) {
			finishedWork := work[wID]
			work[wID] = byte(0)
			res = append(res, finishedWork)
			for _, step := range steps {
				delete(step.before, finishedWork)
			}
		}
	}
	return res, clockTime
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
	order, time := solve(steps)
	fmt.Println(string(order))
	fmt.Println(time)
}
