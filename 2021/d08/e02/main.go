package main

import (
	"adventofcode/utils/files"
	"adventofcode/utils/sets/byteset"
	"os"
	"strconv"
	"strings"
)

//   0:      1:      2:      3:      4:
//  aaaa    ....    aaaa    aaaa    ....
// b    c  .    c  .    c  .    c  b    c
// b    c  .    c  .    c  .    c  b    c
//  ....    ....    dddd    dddd    dddd
// e    f  .    f  e    .  .    f  .    f
// e    f  .    f  e    .  .    f  .    f
//  gggg    ....    gggg    gggg    ....
//
//   5:      6:      7:      8:      9:
//  aaaa    aaaa    aaaa    aaaa    aaaa
// b    .  b    .  .    c  b    c  b    c
// b    .  b    .  .    c  b    c  b    c
//  dddd    dddd    ....    dddd    dddd
// .    f  e    f  .    f  e    f  .    f
// .    f  e    f  .    f  e    f  .    f
//  gggg    gggg    ....    gggg    gggg

var digits = map[int]*byteset.ByteSet{
	1: byteset.NewFromSlice([]byte{'c', 'f'}),
	7: byteset.NewFromSlice([]byte{'a', 'c', 'f'}),
	4: byteset.NewFromSlice([]byte{'b', 'c', 'd', 'f'}),
	2: byteset.NewFromSlice([]byte{'a', 'c', 'd', 'e', 'g'}),
	3: byteset.NewFromSlice([]byte{'a', 'c', 'd', 'f', 'g'}),
	5: byteset.NewFromSlice([]byte{'a', 'b', 'd', 'f', 'g'}),
	0: byteset.NewFromSlice([]byte{'a', 'b', 'c', 'e', 'f', 'g'}),
	6: byteset.NewFromSlice([]byte{'a', 'b', 'd', 'e', 'f', 'g'}),
	9: byteset.NewFromSlice([]byte{'a', 'b', 'c', 'd', 'f', 'g'}),
	8: byteset.NewFromSlice([]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'}),
}

// c = member of 1 that is not a member of one of 0, 6, 9
// f = byte from 1 that is not c
// a = byte from 7 that is not in 1
// g = byte that is in 2, 3, 5, 0, 6, 9 and is not in 7
// d = byte that is in 2, 3, 5 and is not a or g
// b = byte in 4 that is not c, d or f
// e = byte from 8 that is not any of the other segment

// c = member of 1 that is not a member of one of 0, 6, 9
func findC(mapping map[byte]byte, candidates map[int][]*byteset.ByteSet) byte {
	var res byte
	candidates[1][0].Each(func(b byte) {
		matches := 0
		for _, c := range candidates[0] {
			if c.HasMember(b) {
				matches++
			}
		}
		if matches == 2 {
			res = b
		}
	})
	return res
}

// f = byte from 1 that is not c
func findF(mapping map[byte]byte, candidates map[int][]*byteset.ByteSet) byte {
	var res byte
	candidates[1][0].Each(func(b byte) {
		if b != mapping['c'] {
			res = b
		}
	})
	return res
}

// a = byte from 7 that is not in 1
func findA(mapping map[byte]byte, candidates map[int][]*byteset.ByteSet) byte {
	return candidates[7][0].Substract(candidates[1][0]).Members()[0]
}

// g = byte that is in 2, 3, 5, 0, 6, 9 and is not in 7
func findG(mapping map[byte]byte, candidates map[int][]*byteset.ByteSet) byte {
	baseSet := candidates[2][0]
	for _, c := range candidates[2] {
		baseSet = baseSet.Intersection(c)
	}
	for _, c := range candidates[0] {
		baseSet = baseSet.Intersection(c)
	}
	return baseSet.Substract(candidates[7][0]).Members()[0]
}

// d = byte that is in 2, 3, 5 and is not a or g
func findD(mapping map[byte]byte, candidates map[int][]*byteset.ByteSet) byte {
	baseSet := candidates[2][0]
	for _, c := range candidates[2] {
		baseSet = baseSet.Intersection(c)
	}
	var res byte
	baseSet.Each(func(b byte) {
		if b != mapping['a'] && b != mapping['g'] {
			res = b
		}
	})
	return res
}

// b = byte in 4 that is not b, d or f
func findB(mapping map[byte]byte, candidates map[int][]*byteset.ByteSet) byte {
	var res byte
	candidates[4][0].Each(func(c byte) {
		if c != mapping['c'] && c != mapping['d'] && c != mapping['f'] {
			res = c
		}
	})
	return res
}

// e = byte from 8 that is not yet defined
func findE(mapping map[byte]byte, candidates map[int][]*byteset.ByteSet) byte {
	var res byte
	reverseMap := make(map[byte]byte)
	for k, v := range mapping {
		reverseMap[v] = k
	}
	candidates[8][0].Each(func(c byte) {
		if _, found := reverseMap[c]; !found {
			res = c
		}
	})
	return res
}

func mapSegments(input []*byteset.ByteSet) map[byte]byte {
	res := make(map[byte]byte)
	candidates := make(map[int][]*byteset.ByteSet)
	for _, in := range input {
		if in.Len() == digits[1].Len() {
			candidates[1] = append(candidates[1], in)
		} else if in.Len() == digits[7].Len() {
			candidates[7] = append(candidates[7], in)
		} else if in.Len() == digits[4].Len() {
			candidates[4] = append(candidates[4], in)
		} else if in.Len() == digits[8].Len() {
			candidates[8] = append(candidates[8], in)
		} else if in.Len() == digits[2].Len() {
			candidates[2] = append(candidates[2], in)
		} else {
			candidates[0] = append(candidates[0], in)
		}
	}
	res['c'] = findC(res, candidates)
	res['f'] = findF(res, candidates)
	res['a'] = findA(res, candidates)
	res['g'] = findG(res, candidates)
	res['d'] = findD(res, candidates)
	res['b'] = findB(res, candidates)
	res['e'] = findE(res, candidates)
	return res
}

func solve(mapping map[byte]byte, output []*byteset.ByteSet) []byte {
	var res []byte
	reverseMap := make(map[byte]byte)
	for k, v := range mapping {
		reverseMap[v] = k
	}
	for _, out := range output {
		set := byteset.New()
		out.Each(func(b byte) {
			set.Add(reverseMap[b])
		})
		for k, v := range digits {
			if set.Equals(v) {
				res = append(res, byte(k+'0'))
			}
		}
	}
	return res
}

func parse(line string) ([]*byteset.ByteSet, []*byteset.ByteSet) {
	var input []*byteset.ByteSet
	var output []*byteset.ByteSet

	parts := strings.Fields(line)
	for _, p := range parts[0:10] {
		input = append(input, byteset.NewFromSlice([]byte(p)))
	}
	for _, p := range parts[11:] {
		output = append(output, byteset.NewFromSlice([]byte(p)))
	}

	return input, output
}

func atoi(s []byte) int {
	n, err := strconv.Atoi(string(s))
	if err != nil {
		panic(err)
	}
	return n
}

func main() {
	data := files.ReadLines(os.Args[1])
	res := 0
	for _, line := range data {
		input, output := parse(line)
		segments := mapSegments(input)
		res += atoi(solve(segments, output))
	}
}
