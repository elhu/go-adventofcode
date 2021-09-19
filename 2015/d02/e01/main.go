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

func paperReq(box Box) int {
	a := box.l * box.w
	b := box.w * box.h
	c := box.l * box.h

	return 2*a + 2*b + 2*c + min(a, b, c)
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	total := 0
	for _, l := range input {
		b := Box{}
		fmt.Sscanf(l, "%dx%dx%d", &b.l, &b.w, &b.h)
		total += paperReq(b)
	}
	fmt.Println(total)
}
