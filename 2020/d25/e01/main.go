package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func atoi(str string) int {
	i, err := strconv.Atoi(str)
	check(err)
	return i
}

func calcKey(sn int) func() int {
	value := 1
	return func() int {
		value = (value * sn) % 20201227
		return value
	}
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	card := atoi(input[0])
	door := atoi(input[1])
	generator := calcKey(7)
	var cls, dls int
	for i := 1; i < 100000000; i++ {
		k := generator()
		if k == card {
			cls = i
		}
		if k == door {
			dls = i
		}
		if cls != 0 && dls != 0 {
			break
		}
	}
	generator = calcKey(door)
	var encryptionKey int
	for i := 1; i <= cls; i++ {
		encryptionKey = generator()
	}
	fmt.Println(encryptionKey)
}
