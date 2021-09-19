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
	score             int
}

const rounds = 2503

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parse(input []string) map[string]*Reindeer {
	res := make(map[string]*Reindeer)
	for _, l := range input {
		var r Reindeer
		fmt.Sscanf(l, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.", &r.name, &r.speed, &r.burstDuration, &r.restDuration)
		r.stop = r.burstDuration
		res[r.name] = &r
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
		max := 0
		var maxList []string
		for name, r := range reindeers {
			if r.distanceTravelled == max {
				maxList = append(maxList, name)
			}
			if r.distanceTravelled > max {
				max = r.distanceTravelled
				maxList = []string{name}
			}
		}
		for _, name := range maxList {
			reindeers[name].score++
		}
	}
	max := 0
	for _, r := range reindeers {
		if r.score > max {
			max = r.score
		}
	}
	fmt.Println(max)
}
