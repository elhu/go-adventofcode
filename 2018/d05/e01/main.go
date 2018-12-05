package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func kaboom(a, b byte) bool {
	return byte(unicode.ToUpper(rune(a))) == byte(unicode.ToUpper(rune(b))) && a != b
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
