package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Reindeer struct {
	name              string
	speed             int
	burstDuration     int
	restDuration      int
	start             int
	stop              int
	distanceTravelled int
}

const rounds = 2503

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parse(input []string) []*Reindeer {
	var res []*Reindeer
	for _, l := range input {
		var r Reindeer
		fmt.Sscanf(l, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.", &r.name, &r.speed, &r.burstDuration, &r.restDuration)
		r.stop = r.burstDuration
		res = append(res, &r)
	}
	return res
}

func main() {
	rawData, err := ioutil.ReadFile(os.Args[1])
	check(err)
	input := strings.Split(strings.TrimRight(string(rawData), "\n"), "\n")
	reindeers := parse(input)

	for i := 0; i < rounds; i++ {
		for _, r := range reindeers {
			if i >= r.start && i < r.stop {
				r.distanceTravelled += r.speed
			}
			if i == r.stop {
				r.start = i + r.restDuration
				r.stop = r.start + r.burstDuration
			}
		}
	}
	max := 0
	for _, r := range reindeers {
		if r.distanceTravelled > max {
			max = r.distanceTravelled
		}
	}
	fmt.Println(max)
}
