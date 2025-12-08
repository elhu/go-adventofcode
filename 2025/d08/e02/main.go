package main

import (
	"adventofcode/utils/coords/coords3d"
	"adventofcode/utils/files"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

func distance(a, b coords3d.Coords3d) float64 {
	return math.Sqrt(float64((a.X-b.X)*(a.X-b.X) + (a.Y-b.Y)*(a.Y-b.Y) + (a.Z-b.Z)*(a.Z-b.Z)))
}

type DistSortable struct {
	distance float64
	from     coords3d.Coords3d
	to       coords3d.Coords3d
}

func singleCluster(clustersMemberships map[coords3d.Coords3d]coords3d.Coords3d) bool {
	clusterSizes := make(map[coords3d.Coords3d]int)
	for _, cluster := range clustersMemberships {
		clusterSizes[cluster] += 1
	}
	return len(clusterSizes) == 1
}

func solve(dsts []DistSortable) int {
	clustersMemberships := make(map[coords3d.Coords3d]coords3d.Coords3d)
	for _, dst := range dsts {
		clustersMemberships[dst.from] = dst.from
		clustersMemberships[dst.to] = dst.to
	}
	for _, dst := range dsts {
		sourceCluster := clustersMemberships[dst.to]
		targetCluster := clustersMemberships[dst.from]
		if sourceCluster == targetCluster {
			continue
		}
		for member, cluster := range clustersMemberships {
			if cluster == sourceCluster {
				clustersMemberships[member] = targetCluster
			}
		}
		if singleCluster(clustersMemberships) {
			return dst.from.X * dst.to.X
		}
	}
	panic("No single cluster formed")
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	lines := strings.Split(data, "\n")
	var dsts []DistSortable
	var distances []coords3d.Coords3d
	for _, line := range lines {
		var coord coords3d.Coords3d
		fmt.Sscanf(line, "%d,%d,%d", &coord.X, &coord.Y, &coord.Z)
		distances = append(distances, coord)
	}
	for i := 0; i < len(distances)-1; i++ {
		for j := i + 1; j < len(distances); j++ {
			from := distances[i]
			to := distances[j]
			dist := distance(from, to)
			dsts = append(dsts, DistSortable{distance: dist, from: from, to: to})
		}
	}
	sort.Slice(dsts, func(i, j int) bool {
		return dsts[i].distance < dsts[j].distance
	})
	fmt.Println(solve(dsts))
}
