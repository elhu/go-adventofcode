package main

import (
	"adventofcode/utils/files"
	"container/list"
	"fmt"
	"os"
)

func parsePairs(data []string) map[string]byte {
	res := make(map[string]byte)
	var pair string
	var insert byte
	for _, d := range data {
		fmt.Sscanf(d, "%s -> %c", &pair, &insert)
		res[pair] = insert
	}
	return res
}

func parseTemplate(ts string) *list.List {
	list := list.New()
	for _, c := range []byte(ts) {
		list.PushBack(c)
	}
	return list
}

func computeScore(template *list.List) int {
	counts := make(map[byte]int)
	for e := template.Front(); e != nil; e = e.Next() {
		counts[e.Value.(byte)]++
	}
	min := 9999999999
	max := 0
	for _, v := range counts {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return max - min
}

func printList(l *list.List) {
	var data []byte
	for e := l.Front(); e != nil; e = e.Next() {
		data = append(data, e.Value.(byte))
	}
	fmt.Println(string(data))
}

func solve(template *list.List, pairs map[string]byte) int {
	for i := 0; i < turns; i++ {
		for e := template.Front(); e != template.Back(); e = e.Next() {
			n := e.Next()
			k := fmt.Sprintf("%c%c", e.Value.(byte), n.Value.(byte))
			if insert, found := pairs[k]; found {
				template.InsertAfter(insert, e)
				e = e.Next()
			}
		}
	}
	return computeScore(template)
}

const turns = 10

func main() {
	data := files.ReadLines(os.Args[1])
	template := parseTemplate(data[0])
	pairs := parsePairs(data[2:])
	fmt.Println(solve(template, pairs))
}
