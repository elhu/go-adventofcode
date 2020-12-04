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

var mandatoryFields = []string{
	"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid",
}

var validations = map[string](func(string) bool){
	"byr": validateByr,
	"iyr": validateIyr,
	"eyr": validateEyr,
	"hgt": validateHgt,
	"hcl": validateHcl,
	"ecl": validateEcl,
	"pid": validatePid,
	"cid": func(val string) bool { return true },
}

func unsafeAtoi(str string) int {
	val, err := strconv.Atoi(str)
	check(err)
	return val
}

func validateByr(val string) bool {
	n := unsafeAtoi(val)
	return n >= 1920 && n <= 2002
}

func validateIyr(val string) bool {
	n := unsafeAtoi(val)
	return n >= 2010 && n <= 2020
}

func validateEyr(val string) bool {
	n := unsafeAtoi(val)
	return n >= 2020 && n <= 2030
}

func validateHgt(val string) bool {
	if len(val) < 4 {
		return false
	}
	unit := string(val[len(val)-2:])
	n := unsafeAtoi(val[:len(val)-2])
	if unit == "cm" {
		return n >= 150 && n <= 193
	} else if unit == "in" {
		return n >= 59 && n <= 76
	}
	return false
}

func validateHcl(val string) bool {
	if len(val) != 7 || val[0] != '#' {
		return false
	}

	for _, c := range val[1:] {
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f')) {
			return false
		}
	}
	return true
}

func validateEcl(val string) bool {
	validColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, v := range validColors {
		if val == v {
			return true
		}
	}
	return false
}

func validatePid(val string) bool {
	if len(val) != 9 {
		return false
	}
	for _, c := range val {
		if !(c >= '0' && c <= '9') {
			return false
		}
	}
	return true
}

func valid(fields map[string]string) bool {
	for _, m := range mandatoryFields {
		value, present := fields[m]
		// fmt.Println(m, value, present, validations[m](value))
		if !present || !validations[m](value) {
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
