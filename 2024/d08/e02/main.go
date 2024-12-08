package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
	"adventofcode/utils/sets"
	"fmt"
	"os"
	"strings"
)

func findAntennas(lines []string) map[rune]*sets.Set[coords2d.Coords2d] {
	antennas := make(map[rune]*sets.Set[coords2d.Coords2d])
	for y, line := range lines {
		for x, char := range line {
			if char != '.' {
				if _, found := antennas[char]; !found {
					antennas[char] = sets.New[coords2d.Coords2d]()
				}
				antennas[char].Add(coords2d.Coords2d{x, y})
			}
		}
	}
	return antennas
}

func sub(a, b coords2d.Coords2d) coords2d.Coords2d {
	return coords2d.Coords2d{a.X - b.X, a.Y - b.Y}
}

func getAntinodes(a, b coords2d.Coords2d, maxX, maxY int) []coords2d.Coords2d {
	vec := sub(a, b)
	var res []coords2d.Coords2d
	for {
		newPos := coords2d.Add(a, vec)
		if newPos.X >= 0 && newPos.X < maxX && newPos.Y >= 0 && newPos.Y < maxY {
			res = append(res, newPos)
			a = newPos
		} else {
			break
		}
	}
	return res
}

func solve(antennas map[rune]*sets.Set[coords2d.Coords2d], maxX, maxY int) int {
	antinodes := sets.New[coords2d.Coords2d]()
	for _, list := range antennas {
		if list.Len() == 1 {
			continue
		}
		members := list.Members()
		for i, a := range members {
			antinodes.Add(a)
			for _, b := range members[i+1:] {
				for _, an := range getAntinodes(a, b, maxX, maxY) {
					antinodes.Add(an)
				}
				for _, an := range getAntinodes(b, a, maxX, maxY) {
					antinodes.Add(an)
				}
			}
		}
	}
	return antinodes.Len()
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	lines := strings.Split(data, "\n")
	antennas := findAntennas(lines)
	fmt.Println(solve(antennas, len(lines[0]), len(lines)))
}
