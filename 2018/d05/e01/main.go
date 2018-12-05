package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const diff = 'A' - 'a'

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func kaboom(a, b byte) bool {
	return abs(int(a-b)) == diff
}

func reduce(units []byte) ([]byte, int) {
	removed := 0
	res := make([]byte, 0, len(units))
	for i := 0; i < len(units); i++ {
		if i+1 < len(units) && kaboom(units[i], units[i+1]) {
			removed++
			i++
		} else {
			res = append(res, units[i])
		}
	}
	return res, removed
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	data = bytes.TrimSuffix(data, []byte{'\n'})
	check(err)
	for removed := 1; removed != 0; {
		data, removed = reduce(data)
	}
	fmt.Println(len(data))
}
