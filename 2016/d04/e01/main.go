package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var roomExp = regexp.MustCompile(`(?P<name>[\w-]+)-(?P<id>\d+)\[(?P<checksum>\w+)\]`)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readLines(fh *bufio.Reader, c chan string) {
	for {
		line, err := fh.ReadString('\n')
		c <- strings.Trim(line, " \n")
		if err == io.EOF {
			break
		}
	}
	close(c)
}

type room struct {
	name     []byte
	id       int
	checksum []byte
}

func (r *room) isReal() bool {
	freqs := buildFrequencyMap(r.name)
	return bytes.Equal(r.checksum, computeCheckSum(freqs))
}

func computeCheckSum(freqs map[byte]int) []byte {
	type kv struct {
		k byte
		v int
	}
	kvs := make([]kv, 0, len(freqs))
	for k, v := range freqs {
		kvs = append(kvs, kv{k, v})
	}

	sort.Slice(kvs, func(i, j int) bool {
		if kvs[i].v == kvs[j].v {
			return kvs[i].k < kvs[j].k
		}
		return kvs[i].v > kvs[j].v
	})
	res := make([]byte, 5)
	for i, kv := range kvs[0:5] {
		res[i] = kv.k
	}
	return res
}

func buildFrequencyMap(name []byte) map[byte]int {
	freqs := make(map[byte]int)
	for _, b := range name {
		if b != '-' {
			freqs[b]++
		}
	}
	return freqs
}

func parseLine(line string) room {
	match := roomExp.FindStringSubmatch(line)
	name := match[1]
	id, _ := strconv.Atoi(match[2])
	checksum := match[3]
	return room{[]byte(name), id, []byte(checksum)}
}

func main() {
	fh, err := os.Open(os.Args[1])
	check(err)
	defer fh.Close()

	reader := bufio.NewReader(fh)
	c := make(chan string, 100)

	go readLines(reader, c)

	result := 0

	for line := range c {
		if line == "" {
			break
		}
		room := parseLine(line)
		if room.isReal() {
			result += room.id
		}
	}
	fmt.Println(result)
}
