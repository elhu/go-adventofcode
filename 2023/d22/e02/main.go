package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/coords/coords3d"
	"adventofcode/utils/files"
	"adventofcode/utils/sets/intset"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Brick struct {
	id         int
	start, end coords3d.Coords3d
}

func parseBrick(id int, line string) Brick {
	b := Brick{id: id}
	fmt.Sscanf(line, "%d,%d,%d~%d,%d,%d", &b.start.X, &b.start.Y, &b.start.Z, &b.end.X, &b.end.Y, &b.end.Z)
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func land(bricks []Brick) (map[coords3d.Coords3d]int, map[int][]coords3d.Coords3d) {
	coordToBrick := make(map[coords3d.Coords3d]int)    // coords3d.Coords3d -> brick id
	maxZ := make(map[coords2d.Coords2d]int)            // max z for each x,y
	brickToCoords := make(map[int][]coords3d.Coords3d) // Keep a reference of coordinates occupied by each brick
	for _, brick := range bricks {
		mz := maxZ[coords2d.Coords2d{X: brick.start.X, Y: brick.start.Y}]
		for x := brick.start.X; x <= brick.end.X; x++ { // Find the highest z for each x,y, we'll need to rest on top of it
			for y := brick.start.Y; y <= brick.end.Y; y++ {
				mz = max(maxZ[coords2d.Coords2d{X: x, Y: y}], mz)
			}
		}
		// Land on top of the highest z
		for x := brick.start.X; x <= brick.end.X; x++ {
			for y := brick.start.Y; y <= brick.end.Y; y++ {
				for z := 0; z <= brick.end.Z-brick.start.Z; z++ {
					newCoords := coords3d.Coords3d{X: x, Y: y, Z: z + mz + 1}
					coordToBrick[newCoords] = brick.id
					maxZ[coords2d.Coords2d{X: x, Y: y}] = newCoords.Z
					brickToCoords[brick.id] = append(brickToCoords[brick.id], newCoords)
				}
			}
		}
	}
	return coordToBrick, brickToCoords
}

func getSupportedBricks(coordToBrick map[coords3d.Coords3d]int, brickToCoords map[int][]coords3d.Coords3d, brickID int, brokenBricks *intset.IntSet) []int {
	supporting := intset.New()
	for _, coords := range brickToCoords[brickID] {
		if val, found := coordToBrick[coords3d.Coords3d{X: coords.X, Y: coords.Y, Z: coords.Z + 1}]; found && val != brickID {
			supporting.Add(val)
		}
	}
	var res []int
	for _, sID := range supporting.Members() {
		supportedBy := intset.New()
		for _, coords := range brickToCoords[sID] {
			if val, found := coordToBrick[coords3d.Coords3d{X: coords.X, Y: coords.Y, Z: coords.Z - 1}]; found && val != sID && !brokenBricks.HasMember(val) {
				supportedBy.Add(val)
			}
		}
		if supportedBy.Len() == 0 {
			res = append(res, sID)
		}
	}
	return res
}

func explode(coordToBrick map[coords3d.Coords3d]int, brickToCoords map[int][]coords3d.Coords3d, brickID int, brokenBricks *intset.IntSet) {
	brokenBricks.Add(brickID)

	for _, sb := range getSupportedBricks(coordToBrick, brickToCoords, brickID, brokenBricks) {
		explode(coordToBrick, brickToCoords, sb, brokenBricks)
	}
}

func solve(bricks []Brick) int {
	coordToBrick, brickToCoord := land(bricks)
	res := 0

	for _, brick := range bricks {
		brokenBricks := intset.New()
		explode(coordToBrick, brickToCoord, brick.id, brokenBricks)
		res += brokenBricks.Len() - 1
	}
	return res
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	lines := strings.Split(data, "\n")
	var bricks []Brick
	for i, line := range lines {
		bricks = append(bricks, parseBrick(i, line))
	}
	sort.Slice(bricks, func(i, j int) bool { return bricks[i].start.Z < bricks[j].start.Z })
	fmt.Println(solve(bricks))
}
