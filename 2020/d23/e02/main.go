package main

import (
	"bytes"
	"container/ring"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func atoi(str string) int {
	i, err := strconv.Atoi(str)
	check(err)
	return i
}

func elemDo(start *ring.Ring, fn func(*ring.Ring)) {
	e := start
	for ; e.Next() != start; e = e.Next() {
		fn(e)
	}
	fn(e)
}

func inList(i int, list []int) bool {
	for _, d := range list {
		if d == i {
			return true
		}
	}
	return false
}

func findTarget(cupsRef map[int]*ring.Ring, currentValue int, excludedList []int) *ring.Ring {
	for i := currentValue - 1; i > 0; i-- {
		if !inList(i, excludedList) {
			return cupsRef[i]
		}
	}
	for i := 1000000; i > 0; i-- {
		if !inList(i, excludedList) {
			return cupsRef[i]
		}
	}
	panic("WTF")
}

func cupsToString(r *ring.Ring) string {
	res := make([]int, 0)
	s := r
	for ; s.Next() != r; s = s.Next() {
		res = append(res, s.Value.(int))
	}
	res = append(res, s.Value.(int))
	return intSliceJoin(res, ", ")
}

func solve(cupsRef map[int]*ring.Ring, startVal int, turns int) {
	t := cupsRef[startVal]
	for i := 0; i < turns; i++ {
		sub := t.Unlink(3)
		excludedList := make([]int, 0, 3)
		sub.Do(func(val interface{}) {
			excludedList = append(excludedList, val.(int))
		})
		s := findTarget(cupsRef, t.Value.(int), excludedList)
		s.Link(sub)
		t = t.Next()
	}
}

func intSliceJoin(slice []int, sep string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(slice), " ", sep, -1), "[]")
}

func computeResult(cupsRef map[int]*ring.Ring) int {
	t := cupsRef[1]
	return t.Next().Value.(int) * t.Next().Next().Value.(int)
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	data = bytes.TrimRight(data, "\n")
	values := make([]int, len(data))
	for i, d := range data {
		values[i] = atoi(string(d))
	}
	cupsRef := make(map[int]*ring.Ring)
	r := ring.New(1000000)
	for _, d := range values {
		r.Value = d
		cupsRef[d] = r
		r = r.Next()
	}
	for i := 10; i <= 1000000; i++ {
		r.Value = i
		cupsRef[i] = r
		r = r.Next()
	}
	solve(cupsRef, r.Value.(int), 10000000)
	fmt.Println(computeResult(cupsRef))
}
