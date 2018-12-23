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

func botsInRange(bots []*Bot, center *Bot) []*Bot {
	res := make([]*Bot, 0)
	for _, b := range bots {
		if intersect(center, b) {
			res = append(res, b)
		}
	}
	return res
}

func subdivide(center *Bot) []*Bot {
	res := make([]*Bot, 0, 8)
	if center.radius == 0 {
		return res
	}
	newRadius := center.radius / 2
	res = append(res, &Bot{pos: &Position{x: center.pos.x + newRadius, y: center.pos.y + newRadius, z: center.pos.z + newRadius}, radius: newRadius})
	res = append(res, &Bot{pos: &Position{x: center.pos.x + newRadius, y: center.pos.y + newRadius, z: center.pos.z - newRadius}, radius: newRadius})
	res = append(res, &Bot{pos: &Position{x: center.pos.x + newRadius, y: center.pos.y - newRadius, z: center.pos.z + newRadius}, radius: newRadius})
	res = append(res, &Bot{pos: &Position{x: center.pos.x - newRadius, y: center.pos.y + newRadius, z: center.pos.z + newRadius}, radius: newRadius})
	res = append(res, &Bot{pos: &Position{x: center.pos.x + newRadius, y: center.pos.y - newRadius, z: center.pos.z - newRadius}, radius: newRadius})
	res = append(res, &Bot{pos: &Position{x: center.pos.x - newRadius, y: center.pos.y + newRadius, z: center.pos.z - newRadius}, radius: newRadius})
	res = append(res, &Bot{pos: &Position{x: center.pos.x - newRadius, y: center.pos.y - newRadius, z: center.pos.z + newRadius}, radius: newRadius})
	res = append(res, &Bot{pos: &Position{x: center.pos.x - newRadius, y: center.pos.y - newRadius, z: center.pos.z - newRadius}, radius: newRadius})
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func solve3(bots []*Bot) int {
	// center := findMinMaxPosBot(bots)
	// Use output from solve2 to bootstrap to a decent value
	center := &Bot{pos: &Position{x: 53538900, y: 72888900, z: 61988900}}
	center.radius = 0
	matchesAtCenter := len(botsInRange(bots, center))
	origin := &Bot{pos: &Position{x: 0, y: 0, z: 0}}
	distanceFromCenter := distance(center, origin)
	changed := true
	fmt.Printf("Best: %d, center: %v\n", matchesAtCenter, center.pos)
	// fmt.Printf("Starting from %v (%d matches)\n", center.pos, matchesAtCenter)
	newCenter := center
	for changed {
		changed = false
		center = newCenter
		for x := center.pos.x - 1; x <= center.pos.x+1; x++ {
			for y := center.pos.y - 1; y <= center.pos.y+1; y++ {
				for z := center.pos.z - 1; z <= center.pos.z+1; z++ {
					point := &Bot{pos: &Position{x: x, y: y, z: z}}
					// fmt.Printf("Checking %v\n", point.pos)
					matchCount := len(botsInRange(bots, point))
					dist := distance(origin, point)
					if matchCount > matchesAtCenter {
						fmt.Printf("Best: %d, center: %v\n", matchesAtCenter, center.pos)
						newCenter = point
						matchesAtCenter = matchCount
						distanceFromCenter = dist
						changed = true
					} else if matchCount == matchesAtCenter && dist < distanceFromCenter {
						newCenter = point
						distanceFromCenter = dist
						changed = true
					}
				}
			}
		}
	}
	return distanceFromCenter
}

const explorationSize = 100

func solve2(bots []*Bot) int {
	steps := []int{100000, 10000, 1000, 100, 10, 1}
	// steps := []int{10, 1}
	center := findMinMaxPosBot(bots)
	// center := &Bot{pos: &Position{x: 0, y: 0, z: 0}}
	var minDistanceFromCenter int
	for _, ratio := range steps {
		fmt.Printf("Starting from center %v\n", center.pos)
		scaledBots := make([]*Bot, len(bots))
		for i, b := range bots {
			scaledBots[i] = &Bot{pos: &Position{x: b.pos.x / ratio, y: b.pos.y / ratio, z: b.pos.z / ratio}, radius: max(b.radius/ratio, 1)}
		}
		scaledCenter := &Bot{pos: &Position{x: center.pos.x / ratio, y: center.pos.y / ratio, z: center.pos.z / ratio}, radius: 0}
		// fmt.Printf("Checking at scale 1/%d, with center (%v) scaled center %v\n", ratio, center.pos, scaledCenter.pos)
		minDistanceFromCenter = maxInt
		maxCount := 0
		for i := scaledCenter.pos.x - explorationSize; i <= scaledCenter.pos.x+explorationSize; i++ {
			// fmt.Printf("Scale: %d, %d/%d\n", ratio, i, scaledCenter.pos.x+explorationSize)
			for j := scaledCenter.pos.y - explorationSize; j <= scaledCenter.pos.y+explorationSize; j++ {
				for k := scaledCenter.pos.z - explorationSize; k <= scaledCenter.pos.z+explorationSize; k++ {
					point := &Bot{pos: &Position{x: i, y: j, z: k}, radius: 0}
					inRange := botsInRange(scaledBots, point)
					// fmt.Println(len(inRange))
					dist := distance(&Bot{pos: &Position{x: 0, y: 0, z: 0}}, point)
					if len(inRange) > maxCount {
						// fmt.Printf("Setting center to %v\n", point.pos)
						center = &Bot{pos: &Position{x: i, y: j, z: k}, radius: 0}
						// fmt.Printf("Setting min distance from center to %d\n", dist)
						minDistanceFromCenter = dist
						maxCount = len(inRange)
					} else if len(inRange) == maxCount && dist < minDistanceFromCenter {
						// fmt.Printf("Setting center to %v\n", point.pos)
						center = &Bot{pos: &Position{x: i, y: j, z: k}, radius: 0}
						// fmt.Printf("Setting min distance from center to %d\n", dist)
						minDistanceFromCenter = dist
						maxCount = len(inRange)
					}
				}
			}
		}
		fmt.Printf("Max count: %d, center: %v\n", maxCount, center.pos)
		center.pos.x *= ratio
		center.pos.y *= ratio
		center.pos.z *= ratio
	}
	return minDistanceFromCenter
}

func solve4(bots []*Bot) int {
	max := 0
	var best *Bot
	for _, b := range bots {
		inRange := botsInRange(bots, b)
		if len(inRange) > max {
			max = len(inRange)
			best = b
		}
	}
	fmt.Printf("Best: %d\n", max)
	fmt.Printf("Best: %v\n", best.pos)
	best.radius = 0
	fmt.Printf("Woot: %d\n", len(botsInRange(bots, best)))
	return 0
}

func solve(bots []*Bot, center *Bot) int {
	queue := make([]*Bot, 0)
	queue = append(queue, center)
	countPerRadius := make(map[int]int)
	candidatesPerRadius := make(map[int][]*Bot)
	toAdd := make([]*Bot, 0)
	var cube *Bot
	candidates := make([]*Bot, 0)
	var oldRadius = center.radius
	for len(queue) > 0 {
		cube, queue = queue[0], queue[1:]
		toAdd = toAdd[:0]
		var newRadius int
		for _, subcube := range subdivide(cube) {
			newRadius = subcube.radius
			if oldRadius != newRadius {
				fmt.Println(newRadius)
				oldRadius = newRadius
			}
			inRange := botsInRange(bots, subcube)
			if len(inRange) > countPerRadius[subcube.radius] {
				countPerRadius[subcube.radius] = len(inRange)
				toAdd = toAdd[:0]
				toAdd = append(toAdd, subcube)
				candidatesPerRadius[subcube.radius] = candidatesPerRadius[subcube.radius][:0]
				candidatesPerRadius[subcube.radius] = append(candidatesPerRadius[subcube.radius], subcube)
				if subcube.radius == 0 {
					candidates = candidates[:0]
					candidates = append(candidates, subcube)
				}
			} else if len(inRange) == countPerRadius[subcube.radius] {
				toAdd = append(toAdd, subcube)
				candidatesPerRadius[subcube.radius] = append(candidatesPerRadius[subcube.radius], subcube)
				if subcube.radius == 0 {
					candidates = append(candidates, subcube)
				}
			}
		}
		if len(queue) == 0 {
			queue = append(queue, candidatesPerRadius[newRadius]...)
			candidatesPerRadius[newRadius] = candidatesPerRadius[newRadius][:0]
		}
		// queue = append(queue, toAdd...)
	}
	minDistance := maxInt
	for _, c := range candidates {
		d := distance(c, &Bot{pos: &Position{x: 0, y: 0, z: 0}})
		if d < minDistance {
			minDistance = d
		}
	}
	return minDistance
}

const maxInt = int(^uint(0) >> 1)
const minInt = -maxInt - 1

func intersect(a, b *Bot) bool {
	return distance(a, b) <= a.radius+b.radius
}

func findMinMaxPosBot(bots []*Bot) *Bot {
	max := Position{minInt, minInt, minInt}
	min := Position{maxInt, maxInt, maxInt}

	for _, b := range bots {
		if b.pos.x > max.x {
			max.x = b.pos.x
		}
		if b.pos.y > max.y {
			max.y = b.pos.y
		}
		if b.pos.z > max.z {
			max.z = b.pos.z
		}
		if b.pos.x < min.x {
			min.x = b.pos.x
		}
		if b.pos.y < min.y {
			min.y = b.pos.y
		}
		if b.pos.z < min.z {
			min.z = b.pos.z
		}
	}
	center := Position{x: (max.x + min.x) / 2, y: (max.y + min.y) / 2, z: (max.z + min.z) / 2}
	distMax := minInt
	if abs(max.x-min.x) > distMax {
		distMax = abs(max.x - min.x)
	}
	if abs(max.y-min.y) > distMax {
		distMax = abs(max.y - min.y)
	}
	if abs(max.z-min.z) > distMax {
		distMax = abs(max.z - min.z)
	}
	return &Bot{pos: &center, radius: distMax/2 + 1}
}

func main() {
	input, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input = bytes.TrimSuffix(input, []byte{'\n'})
	lines := bytes.Split(input, []byte{'\n'})
	bots := parseBots(lines)
	fmt.Println(solve3(bots))
}
