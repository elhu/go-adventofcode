package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

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

func solve(firstRow string, rowCount int) int {
	rows := make([]string, rowCount)
	rows[0] = firstRow
	for i := 1; i < rowCount; i++ {
		row := bytes.Repeat([]byte{'.'}, len(rows[i-1]))
		for j := 1; j < len(rows[i-1])-1; j++ {
			l, c, r := rows[i-1][j-1], rows[i-1][j], rows[i-1][j+1]
			if (l == '^' && c == '^' && r == '.') ||
				(l == '.' && c == '^' && r == '^') ||
				(l == '^' && c == '.' && r == '.') ||
				(l == '.' && c == '.' && r == '^') {
				row[j] = '^'
			}
		}
		rows[i] = string(row)
	}
	safeCount := 0
	for _, row := range rows {
		// Discount safe spots added for boundary checks
		safeCount += strings.Count(row, ".") - 2
	}
	return safeCount
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	rowCount := atoi(os.Args[2])
	check(err)
	input := strings.TrimRight(string(data), "\n")
	// Add safe spots to avoid checking boundary conditions
	input = fmt.Sprintf(".%s.", input)
	fmt.Println(solve(input, rowCount))
}
