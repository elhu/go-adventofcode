package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strings"
)

func solve(blocks, freeSpaces []int) uint64 {
	freeSpaces = append(freeSpaces, 999999999999)
	pos := blocks[0]
	blocks = blocks[1:]
	pfs := true // Processing Free Spaces
	var res uint64
	blockPos := 0
	for ; ; pos++ {
		if pfs {
			if freeSpaces[0] == 0 {
				freeSpaces = freeSpaces[1:]
				pos--
				pfs = !pfs
			} else {
				res += uint64(pos * len(blocks))
				blocks[len(blocks)-1]--
				freeSpaces[0]--
				if blocks[len(blocks)-1] == 0 {
					blocks = blocks[:len(blocks)-1]
				}
			}
		} else {
			res += uint64(pos * (blockPos + 1))
			blocks[blockPos]--
			if blocks[blockPos] == 0 {
				blockPos++
				pfs = !pfs
			}
		}
		if len(blocks) == 0 || blockPos >= len(blocks) {
			break
		}
	}
	return res
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	var freeSpaces, blocks []int
	for i, c := range data {
		if i%2 == 0 {
			blocks = append(blocks, int(c-'0'))
		} else {
			freeSpaces = append(freeSpaces, int(c-'0'))
		}
	}
	fmt.Println(solve(blocks, freeSpaces))
}
