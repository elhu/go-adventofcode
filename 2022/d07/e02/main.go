package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type File struct {
	name string
	size int
}

type FSNode struct {
	parent   *FSNode
	name     string
	children map[string]*FSNode
	files    map[string]*File
}

func atoi(str string) int {
	res, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return res
}

func parseOutput(data []string) *FSNode {
	root := &FSNode{name: "/", children: make(map[string]*FSNode), files: make(map[string]*File)}
	root.parent = root
	currDir := root
	// Skip first '$ cd /'
	for _, line := range data[1:] {
		parts := strings.Split(line, " ")
		if parts[0] == "$" {
			if parts[1] == "cd" {
				if parts[2] == "/" {
					currDir = root
				} else if parts[2] == ".." {
					currDir = currDir.parent
				} else {
					currDir = currDir.children[parts[2]]
				}
			}
			// Do nothing with ls command
		} else { // Handle output of ls
			if parts[0] == "dir" {
				node, exists := currDir.children[parts[1]]
				if !exists {
					node = &FSNode{name: parts[1], parent: currDir, children: make(map[string]*FSNode), files: make(map[string]*File)}
					currDir.children[parts[1]] = node
				}
			} else {
				size := atoi(parts[0])
				currDir.files[parts[1]] = &File{name: parts[1], size: size}
			}
		}
	}
	return root
}

func totalSize(node *FSNode) int {
	size := 0
	for _, v := range node.files {
		size += v.size
	}
	for _, v := range node.children {
		size += totalSize(v)
	}
	return size
}

func solve(fs *FSNode) int {
	driveSize := 70000000
	usedSpace := totalSize(fs)
	freeSpace := driveSize - usedSpace
	requiredSpace := 30000000
	res := driveSize
	queue := []*FSNode{fs}
	var curr *FSNode
	for len(queue) > 0 {
		curr, queue = queue[0], queue[1:]
		size := totalSize(curr)
		if freeSpace+size >= requiredSpace && size < res {
			res = size
		}
		for _, v := range curr.children {
			queue = append(queue, v)
		}
	}
	return res
}

func main() {
	data := files.ReadLines(os.Args[1])
	fs := parseOutput(data)
	fmt.Println(solve(fs))
}
