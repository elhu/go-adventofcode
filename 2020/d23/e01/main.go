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

func findTarget(r *ring.Ring) *ring.Ring {
	max := r.Value.(int)
	rMax := r
	elemDo(r, func(e *ring.Ring) {
		if e.Value.(int) > max {
			max = e.Value.(int)
			rMax = e
		}
	})
	for i := 1; i <= r.Value.(int); i++ {
		var found *ring.Ring
		elemDo(r, func(e *ring.Ring) {
			if e.Value == r.Value.(int)-i {
				found = e
			}
		})
		if found != nil {
			return found
		}
	}
	return rMax
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

func solve(r *ring.Ring, turns int) {
	t := r
	for i := 0; i < turns; i++ {
		sub := t.Unlink(3)
		s := findTarget(t)
		s.Link(sub)
		t = t.Next()
	}
}

func intSliceJoin(slice []int, sep string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(slice), " ", sep, -1), "[]")
}

func formatResult(r *ring.Ring) string {
	t := r
	for ; t.Next() != r && t.Value.(int) != 1; t = t.Next() {
	}
	t = t.Next()
	res := make([]int, 0)
	for s := t; s.Next() != t; s = s.Next() {
		res = append(res, s.Value.(int))
	}
	return intSliceJoin(res, "")
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	data = bytes.TrimRight(data, "\n")
	values := make([]int, len(data))
	for i, d := range data {
		values[i] = atoi(string(d))
	}
	r := ring.New(len(values))
	for _, d := range values {
		r.Value = d
		r = r.Next()
	}
	solve(r, 100)
	fmt.Println(formatResult(r))
}
