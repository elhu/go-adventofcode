package files

import (
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(filename string) []byte {
	data, err := ioutil.ReadFile(filename)
	check(err)
	return data
}

func ReadLinesWithSeparator(filename string, separator string) []string {
	data := ReadFile(filename)
	return strings.Split(strings.TrimRight(string(data), "\n"), separator)
}

func ReadLines(filename string) []string {
	return ReadLinesWithSeparator(filename, "\n")
}
