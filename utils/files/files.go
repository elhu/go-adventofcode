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

// ReadFile reads the entire file and returns it as byte slice.
// Errors are considered unrecoverable will crash the program.
func ReadFile(filename string) []byte {
	data, err := ioutil.ReadFile(filename)
	check(err)
	return data
}

// ReadLinesWithSeparator reads the entire file, splits it on <separator>
// and returns the result as a slice of strings.
// Errors are considered unrecoverable will crash the program.
func ReadLinesWithSeparator(filename string, separator string) []string {
	data := ReadFile(filename)
	return strings.Split(strings.TrimRight(string(data), separator), separator)
}

// ReadLinesWithSeparator reads the entire file, splits it on newlines
// and returns the result as a slice of strings.
// Errors are considered unrecoverable will crash the program.
func ReadLines(filename string) []string {
	return ReadLinesWithSeparator(filename, "\n")
}
