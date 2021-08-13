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
	c := r
	for i := 0; i < elfCount/2; i++ {
		c = c.Next()
	}
	for r.Next() != r {
		c = c.Next()
		c.Prev().Prev().Unlink(1)
		if elfCount%2 == 1 {
			c = c.Next()
		}
		elfCount--
		r = r.Next()
	}
	return r.Value.(int)
}

func main() {
	elfCount := atoi(os.Args[1])
	fmt.Println(solve(elfCount))
}
