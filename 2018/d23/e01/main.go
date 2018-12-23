package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var botExp = regexp.MustCompile(`pos=<(-?\d+),(-?\d+),(-?\d+)>, r=(-?\d+)`)

type Position struct {
	x, y, z int
}

type Bot struct {
	radius int
	pos    *Position
}

func parseBots(input [][]byte) []*Bot {
	res := make([]*Bot, 0, len(input))
	for _, line := range input {
		match := botExp.FindStringSubmatch(string(line))
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		z, _ := strconv.Atoi(match[3])
		radius, _ := strconv.Atoi(match[4])
		res = append(res, &Bot{radius: radius, pos: &Position{x: x, y: y, z: z}})
	}
	return res
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func distance(a, b *Bot) int {
	return abs(a.pos.x-b.pos.x) + abs(a.pos.y-b.pos.y) + abs(a.pos.z-b.pos.z)
}

func solve(bots []*Bot) int {
	var maxRadiusBot *Bot
	maxRadius := 0
	for _, b := range bots {
		if b.radius > maxRadius {
			maxRadius = b.radius
			maxRadiusBot = b
		}
	}
	res := 0
	for _, b := range bots {
		if distance(maxRadiusBot, b) <= maxRadiusBot.radius {
			res++
		}
	}
	return res
}

func main() {
	input, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input = bytes.TrimSuffix(input, []byte{'\n'})
	lines := bytes.Split(input, []byte{'\n'})
	bots := parseBots(lines)
	fmt.Println(solve(bots))
}
