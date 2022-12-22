package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

var directions = []coords2d.Coords2d{
	{X: 1, Y: 0},  // right
	{X: 0, Y: 1},  // down
	{X: -1, Y: 0}, // left
	{X: 0, Y: -1}, // up
}

func findStartTile(board [][]byte) coords2d.Coords2d {
	for i := 0; i < len(board[0]); i++ {
		if board[0][i] == '.' {
			return coords2d.Coords2d{X: i, Y: 0}
		}
	}
	panic("No starting tile found")
}

func oob(board [][]byte, pos coords2d.Coords2d) bool {
	return pos.Y < 0 || pos.Y >= len(board) || pos.X < 0 || pos.X >= len(board[pos.Y])
}

func wrap(board [][]byte, pos, vec coords2d.Coords2d) coords2d.Coords2d {
	newPos := coords2d.Add(pos, vec)
	newPos.Y = mod(newPos.Y, len(board))
	newPos.X = mod(newPos.X, len(board[newPos.Y]))
	for oob(board, newPos) || board[newPos.Y][newPos.X] == ' ' {
		newPos = coords2d.Add(newPos, vec)
		newPos.Y = mod(newPos.Y, len(board))
		newPos.X = mod(newPos.X, len(board[newPos.Y]))
	}
	return newPos
}

func move(board [][]byte, curr coords2d.Coords2d, dirIndex, distance int) coords2d.Coords2d {
	vec := directions[dirIndex]
	for i := 0; i < distance; i++ {
		next := coords2d.Add(curr, vec)
		if oob(board, next) || board[next.Y][next.X] == ' ' {
			next = wrap(board, curr, vec)
		}
		if board[next.Y][next.X] == '.' {
			curr = next
		} else {
			return curr
		}
	}
	return curr
}

func mod(i, m int) int {
	res := i % m
	if res < 0 {
		return res + m
	}
	return res
}

func solve(board [][]byte, moves []byte) int {
	curr := findStartTile(board)
	dirIndex := 0
	var distance int
	var direction byte
	for len(moves) > 0 {
		fmt.Sscanf(string(moves), "%d", &distance)
		moves = moves[len(strconv.Itoa(distance)):]
		curr = move(board, curr, dirIndex, distance)
		if len(moves) > 0 {
			direction = moves[0]
			if direction == 'L' {
				dirIndex = mod(dirIndex-1, len(directions))
			} else {
				dirIndex = mod(dirIndex+1, len(directions))
			}
			moves = moves[1:]
		}
	}
	return 1000*(curr.Y+1) + 4*(curr.X+1) + dirIndex
}

func padBoard(board [][]byte) [][]byte {
	maxLen := 0
	for _, b := range board {
		if len(b) > maxLen {
			maxLen = len(b)
		}
	}
	res := make([][]byte, len(board))
	for i, b := range board {
		res[i] = make([]byte, maxLen)
		for j := 0; j < maxLen; j++ {
			res[i][j] = ' '
		}
		copy(res[i], b)
	}
	return res
}

func main() {
	data := bytes.TrimRight(files.ReadFile(os.Args[1]), "\n")
	parts := bytes.Split(data, []byte("\n\n"))
	board := bytes.Split(parts[0], []byte("\n"))
	board = padBoard(board)
	moves := parts[len(parts)-1]
	fmt.Println(solve(board, moves))
}
