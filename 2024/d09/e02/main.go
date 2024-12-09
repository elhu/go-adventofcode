package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strings"
)

type blockType int

const (
	FILE blockType = iota
	FREE blockType = iota
)

type block struct {
	id, size  int
	kind      blockType
	processed bool
}

func createBlocks(files, freeSpaces []int) []*block {
	var blocks []*block
	freeSpaces = append(freeSpaces, 0)
	for i := 0; i < len(files); i++ {
		blocks = append(blocks, &block{id: i, size: files[i], kind: FILE, processed: false})
		blocks = append(blocks, &block{size: freeSpaces[i], kind: FREE})
	}
	return blocks
}

func findMatch(size int, blocks []*block) int {
	for idx, b := range blocks {
		if b.kind == FREE && b.size >= size {
			return idx
		}
	}
	return -1
}

func moveBlock(blocks []*block, from, to int) []*block {
	if from < to {
		return blocks
	}
	// Reduce the size of the free space at destination
	blocks[to].size -= blocks[from].size
	// Expand the slice
	blocks = append(blocks[:to+1], blocks[to:]...)
	// Insert the element at the destination
	blocks[to] = blocks[from+1]
	// Replace the element with free space
	blocks[from+1] = &block{size: blocks[to].size, kind: FREE}
	return blocks
}

func defrag(files, freeSpaces []int) []*block {
	blocks := createBlocks(files, freeSpaces)
	for i := len(blocks) - 1; i >= 0; i-- {
		if blocks[i].kind == FILE && !blocks[i].processed {
			blocks[i].processed = true
			if idx := findMatch(blocks[i].size, blocks); idx != -1 {
				blocks = moveBlock(blocks, i, idx)
				i++
			}
		}
	}
	return blocks
}

func solve(files, freeSpaces []int) int {
	blocks := defrag(files, freeSpaces)
	res, pos := 0, 0
	for _, block := range blocks {
		for j := 0; j < block.size; j++ {
			if block.kind == FILE {
				res += pos * block.id
			}
			pos++
		}
	}
	return res
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	var freeSpaces, files []int
	for i, c := range data {
		if i%2 == 0 {
			files = append(files, int(c-'0'))
		} else {
			freeSpaces = append(freeSpaces, int(c-'0'))
		}
	}
	fmt.Println(solve(files, freeSpaces))
}
