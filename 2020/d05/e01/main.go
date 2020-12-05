package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func computeRow(str string) int64 {
	str = strings.ReplaceAll(str, "F", "0")
	str = strings.ReplaceAll(str, "B", "1")
	str = fmt.Sprintf("0b%s", str)
	i, err := strconv.ParseInt(str, 0, 64)
	check(err)
	return i
}

func computeColumn(str string) int64 {
	str = strings.ReplaceAll(str, "L", "0")
	str = strings.ReplaceAll(str, "R", "1")
	str = fmt.Sprintf("0b%s", str)
	i, err := strconv.ParseInt(str, 0, 64)
	check(err)
	return i
}

func computeSeatID(ref string) int64 {
	r := computeRow(ref[:7])
	c := computeColumn(ref[7:])
	return r*8 + c
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	lines := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	max := int64(0)
	for _, l := range lines {
		id := computeSeatID(l)
		if id > max {
			max = id
		}
	}
	fmt.Println(max)
}
