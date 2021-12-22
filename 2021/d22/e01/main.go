package main

import (
	"adventofcode/utils/files"
	"adventofcode/utils/sets/stringset"
	"fmt"
	"os"
)

// func hash2(a, b int) int {
// 	return (a+b)*(a+b+1) + b
// }

// func hash3(a, b, c int) int {
// 	return (hash2(a, b)+c)*(hash2(a, b)+c+1) + c
// }

func hash3(a, b, c int) string {
	return fmt.Sprintf("%d:%d:%d", a, b, c)
}

func buildCuboid(minX, maxX, minY, maxY, minZ, maxZ int) *stringset.StringSet {
	res := stringset.New()
	if minX == maxX || minY == maxY || minZ == maxZ {
		return res
	}
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			for z := minZ; z <= maxZ; z++ {
				res.Add(hash3(x, y, z))
			}
		}
	}
	return res
}

func cap(args ...int) (int, int, int, int, int, int) {
	res := make([]int, 6)
	minCap := args[0]
	maxCap := args[1]
	for i, n := range args[2:] {
		if n < minCap {
			n = minCap
		}
		if n > maxCap {
			n = maxCap
		}
		res[i] = n
	}
	return res[0], res[1], res[2], res[3], res[4], res[5]
}

func solve(input []string) int {
	reactor := stringset.New()
	for _, l := range input {
		var inst string
		var minX, maxX, minY, maxY, minZ, maxZ int
		fmt.Sscanf(l, "%s x=%d..%d,y=%d..%d,z=%d..%d", &inst, &minX, &maxX, &minY, &maxY, &minZ, &maxZ)
		minX, maxX, minY, maxY, minZ, maxZ = cap(-50, 50, minX, maxX, minY, maxY, minZ, maxZ)
		cuboid := buildCuboid(minX, maxX, minY, maxY, minZ, maxZ)
		if cuboid.Len() > 0 {
			switch inst {
			case "on":
				reactor = reactor.Union(cuboid)
			case "off":
				reactor = reactor.Substract(cuboid)
			default:
				panic("WTF")
			}
		}
	}

	return reactor.Len()
}

func main() {
	data := files.ReadLines(os.Args[1])
	fmt.Println(solve(data))
}
