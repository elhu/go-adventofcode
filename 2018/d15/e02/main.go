package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Kind byte

const (
	Wall   Kind = iota
	Empty  Kind = iota
	Goblin Kind = iota
	Elf    Kind = iota
)

const DefaultHitPoints = 200
const DefaultAttack = 3

type coord struct {
	x, y int
}

const keyOffset = 100000

func (c *coord) toKey() int {
	return c.y*keyOffset + c.x
}

type cell struct {
	kind         Kind
	hitPoints    int
	attackPoints int
	played       bool
	pos          *coord
}

type sortCells []*cell

func (s sortCells) Less(i, j int) bool {
	return s[i].pos.toKey() < s[j].pos.toKey()
}

func (s sortCells) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortCells) Len() int {
	return len(s)
}

func parseMap(raw [][]byte, atk int) [][]*cell {
	res := make([][]*cell, len(raw))
	for i, line := range raw {
		res[i] = make([]*cell, len(raw[i]))
		for j, c := range line {
			pos := &coord{x: j, y: i}
			switch c {
			case '#':
				res[i][j] = &cell{kind: Wall, pos: pos}
			case '.':
				res[i][j] = &cell{kind: Empty, pos: pos}
			case 'G':
				res[i][j] = &cell{kind: Goblin, hitPoints: DefaultHitPoints, attackPoints: DefaultAttack, pos: pos}
			case 'E':
				res[i][j] = &cell{kind: Elf, hitPoints: DefaultHitPoints, attackPoints: atk, pos: pos}
			}
		}
	}
	return res
}

func resetPlayed(plan [][]*cell) {
	for i := 0; i < len(plan); i++ {
		for j := 0; j < len(plan[i]); j++ {
			plan[i][j].played = false
		}
	}
}

func (c *cell) attack(plan [][]*cell, target Kind) {
	coords := [4]coord{
		coord{y: c.pos.y - 1, x: c.pos.x},
		coord{y: c.pos.y, x: c.pos.x - 1},
		coord{y: c.pos.y, x: c.pos.x + 1},
		coord{y: c.pos.y + 1, x: c.pos.x},
	}
	targets := make([]*cell, 0, 4)
	for _, pos := range coords {
		if plan[pos.y][pos.x].kind == target {
			targets = append(targets, plan[pos.y][pos.x])
		}
	}
	if len(targets) > 0 {
		minHP := 201
		var minTarget *cell
		for _, t := range targets {
			if t.hitPoints < minHP {
				minTarget = t
				minHP = t.hitPoints
			}
		}
		minTarget.hitPoints -= c.attackPoints
		if minTarget.hitPoints <= 0 {
			if minTarget.kind == Elf {
				elfDied++
			}
			minTarget.kind = Empty
		}
	}
}

func (c *cell) findCandidates(plan [][]*cell, target Kind) []*cell {
	candidates := make([]*cell, 0)
	for _, l := range plan {
		for _, curr := range l {
			if curr.kind == Empty {
				coords := [4]coord{
					coord{y: curr.pos.y - 1, x: curr.pos.x},
					coord{y: curr.pos.y, x: curr.pos.x - 1},
					coord{y: curr.pos.y, x: curr.pos.x + 1},
					coord{y: curr.pos.y + 1, x: curr.pos.x},
				}
				for _, pos := range coords {
					if plan[pos.y][pos.x].kind == target {
						candidates = append(candidates, plan[curr.pos.y][curr.pos.x])
					}
				}
			}
		}
	}
	return candidates
}

func distance(start, dest *cell, plan [][]*cell) int {
	open := make([]*cell, 1)
	distances := make(map[int]int)
	distances[start.pos.toKey()] = 0
	open[0] = start
	var curr *cell
	for len(open) > 0 {
		curr, open = open[0], open[1:]
		coords := [4]coord{
			coord{y: curr.pos.y - 1, x: curr.pos.x},
			coord{y: curr.pos.y, x: curr.pos.x - 1},
			coord{y: curr.pos.y, x: curr.pos.x + 1},
			coord{y: curr.pos.y + 1, x: curr.pos.x},
		}
		for _, pos := range coords {
			i := pos.y
			j := pos.x
			if plan[i][j].kind == Empty || plan[i][j].pos.toKey() == dest.pos.toKey() {
				_, visited := distances[plan[i][j].pos.toKey()]
				if !visited {
					open = append(open, plan[i][j])
					distances[plan[i][j].pos.toKey()] = distances[curr.pos.toKey()] + 1
				}
				if plan[i][j].pos.toKey() == dest.pos.toKey() {
					return distances[curr.pos.toKey()] + 1
				}
			}
		}
	}
	return -1
}

func (c *cell) closest(candidates []*cell, plan [][]*cell) (*cell, bool) {
	var res *cell
	minDist := keyOffset*2 + 1
	found := false
	for _, cand := range candidates {
		if dist := distance(c, cand, plan); dist != -1 {
			found = true
			if dist < minDist {
				res = cand
				minDist = dist
			} else if dist == minDist && cand.pos.toKey() < res.pos.toKey() {
				res = cand
			}
		}
	}
	return res, found
}

func (c *cell) move(plan [][]*cell, target Kind) *coord {
	// Check for targets already in range
	coords := [4]coord{
		coord{y: c.pos.y - 1, x: c.pos.x},
		coord{y: c.pos.y, x: c.pos.x - 1},
		coord{y: c.pos.y, x: c.pos.x + 1},
		coord{y: c.pos.y + 1, x: c.pos.x},
	}
	for _, pos := range coords {
		if plan[pos.y][pos.x].kind == target {
			return c.pos
		}
	}
	// List all candidate targets
	candidates := c.findCandidates(plan, target)
	// Find the one closest to c and first in reading order
	if target, found := c.closest(candidates, plan); found {
		newPos := findFirstStep(c, target, plan)
		if newPos.toKey() != c.pos.toKey() {
			plan[c.pos.y][c.pos.x] = &cell{kind: Empty, pos: &coord{x: c.pos.x, y: c.pos.y}}
			plan[newPos.y][newPos.x] = &cell{
				kind:         c.kind,
				hitPoints:    c.hitPoints,
				attackPoints: c.attackPoints,
				pos:          &coord{x: newPos.x, y: newPos.y},
				played:       true,
			}
			return newPos
		}
	}
	return c.pos
}

func findFirstStep(start, target *cell, plan [][]*cell) *coord {
	open := make([]*cell, 1)
	distances := make(map[int]int)
	distances[target.pos.toKey()] = 0
	open[0] = target
	var curr *cell
	for len(open) > 0 {
		curr, open = open[0], open[1:]
		coords := [4]coord{
			coord{y: curr.pos.y - 1, x: curr.pos.x},
			coord{y: curr.pos.y, x: curr.pos.x - 1},
			coord{y: curr.pos.y, x: curr.pos.x + 1},
			coord{y: curr.pos.y + 1, x: curr.pos.x},
		}
		for _, pos := range coords {
			i := pos.y
			j := pos.x
			if plan[i][j].kind == Empty || plan[i][j].pos.toKey() == start.pos.toKey() {
				_, visited := distances[plan[i][j].pos.toKey()]
				if !visited {
					open = append(open, plan[i][j])
					distances[plan[i][j].pos.toKey()] = distances[curr.pos.toKey()] + 1
				}
			}
		}
	}
	coords := [4]coord{
		coord{y: start.pos.y - 1, x: start.pos.x},
		coord{y: start.pos.y, x: start.pos.x - 1},
		coord{y: start.pos.y, x: start.pos.x + 1},
		coord{y: start.pos.y + 1, x: start.pos.x},
	}
	minDistance := keyOffset*2 + 1
	var res coord
	for _, c := range coords {
		dist, exists := distances[c.toKey()]
		if exists && dist < minDistance {
			minDistance = dist
			res = c
		}
	}
	return &res
}

func playCell(plan [][]*cell, pos coord, target Kind) {
	newPos := plan[pos.y][pos.x].move(plan, target)
	plan[newPos.y][newPos.x].attack(plan, target)
}

func printPlan(plan [][]*cell) {
	for _, l := range plan {
		cells := make([]*cell, 0)
		for _, c := range l {
			switch c.kind {
			case Empty:
				fmt.Printf(".")
			case Wall:
				fmt.Printf("#")
			case Goblin:
				cells = append(cells, c)
				fmt.Printf("G")
			case Elf:
				cells = append(cells, c)
				fmt.Printf("E")
			}
		}
		for _, c := range cells {
			if c.kind == Goblin {
				fmt.Printf(" G(%d)", c.hitPoints)
			} else {
				fmt.Printf(" E(%d)", c.hitPoints)
			}
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func present(plan [][]*cell, target Kind) bool {
	for _, l := range plan {
		for _, c := range l {
			if c.kind == target {
				return true
			}
		}
	}
	return false
}

func play(plan [][]*cell) int {
	rounds := 0
	for present(plan, Goblin) && present(plan, Elf) {
		for i := 0; i < len(plan); i++ {
			for j := 0; j < len(plan[i]); j++ {
				if c := plan[i][j]; !c.played {
					switch c.kind {
					case Goblin:
						if !present(plan, Elf) {
							return rounds
						}
						playCell(plan, coord{x: j, y: i}, Elf)
					case Elf:
						if !present(plan, Goblin) {
							return rounds
						}
						playCell(plan, coord{x: j, y: i}, Goblin)
					}
					plan[i][j].played = true
				}
			}
		}
		rounds++
		resetPlayed(plan)
	}
	return rounds
}

func sumRemainingHP(plan [][]*cell) int {
	res := 0
	for _, l := range plan {
		for _, c := range l {
			switch c.kind {
			case Goblin:
				res += c.hitPoints
			case Elf:
				res += c.hitPoints
			}
		}
	}
	return res
}

var elfDied = 0

func main() {
	input, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input = bytes.TrimSuffix(input, []byte{'\n'})
	rawMap := bytes.Split(input, []byte{'\n'})
	var plan [][]*cell
	var rounds int
	for atk := 4; ; atk++ {
		elfDied = 0
		plan = parseMap(rawMap, atk)
		rounds = play(plan)
		if elfDied == 0 {
			break
		}
	}
	printPlan(plan)
	fmt.Printf("Rounds played: %d\n", rounds)
	fmt.Printf("Remaining HP: %d\n", sumRemainingHP(plan))
	fmt.Println(rounds * sumRemainingHP(plan))
}
