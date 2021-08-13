package main

import (
	"fmt"
	"os"
	"strconv"
)

func atoi(str string) int {
	res, err := strconv.Atoi(str)
	check(err)
	return res
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func computeCheckSum(data []byte) []byte {
	var sum []byte
	for len(sum)%2 != 1 {
		sum = nil
		for i := 0; i < len(data)-1; i += 2 {
			if data[i] == data[i+1] {
				sum = append(sum, '1')
			} else {
				sum = append(sum, '0')
			}
		}
		data = sum
	}
	return sum
}

func solve(seed []byte, diskSize int) []byte {
	res := make([]byte, len(seed))
	copy(res, seed)
	for len(res) < diskSize {
		b := make([]byte, len(res))
		for i := 0; i < len(res); i++ {
			if res[i] == '0' {
				b[len(res)-i-1] = '1'
			} else {
				b[len(res)-i-1] = '0'
			}
		}
		res = append(res, '0')
		res = append(res, b...)
	}
	return computeCheckSum(res[0:diskSize])
}

func main() {
	seed := []byte(os.Args[1])
	diskSize := atoi(os.Args[2])
	fmt.Println(string(solve(seed, diskSize)))
}
