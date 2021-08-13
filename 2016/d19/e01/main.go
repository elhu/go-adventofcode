package main

import (
	"container/ring"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func atoi(str string) int {
	res, err := strconv.Atoi(str)
	check(err)
	return res
}

func solve(elfCount int) int {
	r := ring.New(elfCount)
	for i := 0; i < elfCount; i++ {
		r.Value = i + 1
		r = r.Next()
	}
	for r.Next() != r {
		r.Unlink(1)
		r = r.Next()
	}
	return r.Value.(int)
}

func main() {
	elfCount := atoi(os.Args[1])
	fmt.Println(solve(elfCount))
}
