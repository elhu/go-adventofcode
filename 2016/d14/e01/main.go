package main

import (
	"crypto/md5"
	"fmt"
	"os"
)

func indexHash(salt string, index int) string {
	toHash := []byte(fmt.Sprintf("%s%d", salt, index))
	return fmt.Sprintf("%x", md5.Sum(toHash))
}

func hasQuintuple(c byte, nextK []string) bool {
	for _, h := range nextK {
		for i := 0; i < len(h)-4; i++ {
			if h[i] == c && h[i+1] == c && h[i+2] == c && h[i+3] == c && h[i+4] == c {
				return true
			}
		}
	}
	return false
}

func isKey(hash string, nextK []string) bool {
	for i := 0; i < len(hash)-2; i++ {
		if hash[i] == hash[i+1] && hash[i] == hash[i+2] {
			return hasQuintuple(hash[i], nextK)
		}
	}
	return false
}

func solve(salt string) int {
	var keys []int
	idx := 0
	hashes := make([]string, 1001)
	for delta := 0; delta <= 1000; delta++ {
		hashes[delta] = indexHash(salt, delta)
	}
	for ; len(keys) < 64; idx++ {
		currHash := hashes[0]
		if isKey(currHash, hashes[1:]) {
			keys = append(keys, idx)
		}
		hashes = append(hashes[1:], indexHash(salt, idx+1001))
	}
	return keys[len(keys)-1]
}

func main() {
	salt := os.Args[1]
	fmt.Println(solve(salt))
}
