package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// Perm calls f with each permutation of a.
func Perm(a []byte, f func([]byte)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []byte, f func([]byte), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func atoi(str string) int {
	res, err := strconv.Atoi(str)
	check(err)
	return res
}

func swapPos(pwd []byte, a, b int) []byte {
	pwd[a], pwd[b] = pwd[b], pwd[a]
	return pwd
}

func swapLetter(pwd []byte, a, b byte) []byte {
	for i, c := range pwd {
		if c == a {
			pwd[i] = b
		} else if c == b {
			pwd[i] = a
		}
	}
	return pwd
}

func rotateLeft(pwd []byte, x int) []byte {
	res := append([]byte{}, pwd[x:]...)
	res = append(res, pwd[0:x]...)
	return res
}

func rotateRight(pwd []byte, x int) []byte {
	return rotateLeft(pwd, len(pwd)-x)
}

func rotatePos(pwd []byte, c []byte) []byte {
	idx := bytes.Index(pwd, c)
	if idx >= 4 {
		return rotateRight(pwd, (idx+2)%len(pwd))
	}
	return rotateRight(pwd, (idx+1)%len(pwd))

}

func reverse(pwd []byte, a, b int) []byte {
	for i := 0; i <= (b-a)/2; i++ {
		pwd[a+i], pwd[b-i] = pwd[b-i], pwd[a+i]
	}
	return pwd
}

func move(pwd []byte, a, b int) []byte {
	res := make([]byte, len(pwd))
	target := pwd[a]
	pwd = append(pwd[0:a], pwd[a+1:]...)
	// fmt.Printf("%s\n", pwd)
	copy(res, pwd[0:b])
	res[b] = target
	// res := append(pwd[0:b], target)
	// fmt.Printf("%s + %s\n", res, pwd[b:])
	copy(res[b+1:], pwd[b:])
	// res = append(res, pwd[b:]...)
	return res
}

func scramble(input []string, pwd []byte) []byte {
	for _, l := range input {
		if strings.HasPrefix(l, "swap position") {
			var a, b int
			fmt.Sscanf(l, "swap position %d with position %d", &a, &b)
			pwd = swapPos(pwd, a, b)
		} else if strings.HasPrefix(l, "swap letter") {
			var a, b []byte
			fmt.Sscanf(l, "swap letter %s with letter %s", &a, &b)
			pwd = swapLetter(pwd, a[0], b[0])
		} else if strings.HasPrefix(l, "rotate left") {
			var a int
			fmt.Sscanf(l, "rotate left %d step", &a)
			pwd = rotateLeft(pwd, a)
		} else if strings.HasPrefix(l, "rotate right") {
			var a int
			fmt.Sscanf(l, "rotate right %d step", &a)
			pwd = rotateRight(pwd, a)
		} else if strings.HasPrefix(l, "rotate based") {
			var a []byte
			fmt.Sscanf(l, "rotate based on position of letter %s", &a)
			pwd = rotatePos(pwd, a)
		} else if strings.HasPrefix(l, "reverse") {
			var a, b int
			fmt.Sscanf(l, "reverse positions %d through %d", &a, &b)
			pwd = reverse(pwd, a, b)
		} else if strings.HasPrefix(l, "move position") {
			var a, b int
			fmt.Sscanf(l, "move position %d to position %d", &a, &b)
			pwd = move(pwd, a, b)
		}
	}
	return pwd
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	pwd := []byte(os.Args[2])
	Perm(pwd, func(p []byte) {
		cpy := make([]byte, len(p))
		copy(cpy, p)
		out := scramble(input, cpy)
		// fmt.Printf("[%s] %s => %s\n", pwd, p, string(out))
		if bytes.Equal(out, []byte(os.Args[2])) {
			fmt.Println(string(p))
			return
		}
	})
	// fmt.Println(string())
}
