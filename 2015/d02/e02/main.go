package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Box struct {
	l, w, h int
}

func min(a, b, c int) int {
	min := a
	for _, n := range []int{a, b, c} {
		if n < min {
			min = n
		}
	}
	return min
}

func ribbonReq(box Box) int {
	a := 2*box.l + 2*box.w
	b := 2*box.w + 2*box.h
	c := 2*box.l + 2*box.h
	return box.l*box.w*box.h + min(a, b, c)
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	total := 0
	for _, l := range input {
		b := Box{}
		fmt.Sscanf(l, "%dx%dx%d", &b.l, &b.w, &b.h)
		total += ribbonReq(b)
	}
	fmt.Println(total)
}
