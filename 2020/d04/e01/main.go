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

var mandatoryFields = []string{
	"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid",
}

func valid(fields map[string]string) bool {
	for _, m := range mandatoryFields {
		if _, present := fields[m]; !present {
			return false
		}
	}
	return true
}

func parseIDDoc(fields []string) map[string]string {
	res := make(map[string]string)
	for _, f := range fields {
		parts := strings.Split(f, ":")
		res[parts[0]] = parts[1]
	}
	return res
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	lines := strings.Split(strings.TrimRight(string(data), "\n"), "\n\n")
	res := 0
	for _, l := range lines {
		fields := strings.Fields(l)
		if valid(parseIDDoc(fields)) {
			res++
		}
	}
	fmt.Println(res)
}
