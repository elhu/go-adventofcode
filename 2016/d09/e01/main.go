package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type marker struct {
	strLength    int
	repeatLength int
	repeatTimes  int
}

var markerExp = regexp.MustCompile(`^\((\d+)x(\d+)\)`)

func processMarker(data []byte) marker {

	match := markerExp.FindStringSubmatch(string(data))
	repeatLength, _ := strconv.Atoi(match[1])
	repeatTimes, _ := strconv.Atoi(match[2])
	return marker{len(match[0]), repeatLength, repeatTimes}
}

func decompress(data []byte) int {
	var res int
	var pos int
	for pos = 0; pos < len(data); {
		if data[pos] == '(' {
			mark := processMarker(data[pos:])
			pos += mark.strLength + mark.repeatLength
			res += mark.repeatLength * mark.repeatTimes
		} else {
			pos++
			res++
		}
	}
	return res
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	res := decompress([]byte(strings.Replace(strings.Trim(string(data), " \n"), " ", "", -1)))
	fmt.Println(res)
}
