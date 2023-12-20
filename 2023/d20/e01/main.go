package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strings"
)

const (
	BROADCASTER = "broadcaster"
	BUTTON      = "button"
	FLIP_FLOP   = "flip-flop"
	CONJUNCTION = "conjunction"
	UNKNOWN     = "unknown"
)

const (
	LOW  = false
	HIGH = true
)

const (
	OFF = false
	ON  = true
)

type Signal struct {
	kind     bool
	from, to *Module
}

type Module struct {
	kind          string
	name          string
	sendsTo       []*Module
	receivesFrom  []*Module
	flipFlopState bool
	conjState     map[string]bool
}

func parseModules(lines []string) map[string]*Module {
	modules := make(map[string]*Module)
	button := &Module{kind: BUTTON, name: BUTTON}
	broadcaster := &Module{kind: BROADCASTER, name: BROADCASTER}
	button.sendsTo = append(button.sendsTo, broadcaster)
	broadcaster.receivesFrom = append(broadcaster.receivesFrom, button)
	modules[BUTTON] = button
	modules[BROADCASTER] = broadcaster
	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		m := &Module{}
		if parts[0][0] == '&' {
			m.kind = CONJUNCTION
			m.name = parts[0][1:]
		} else if parts[0][0] == '%' {
			m.kind = FLIP_FLOP
			m.name = parts[0][1:]
		} else if parts[0] != BROADCASTER {
			m.kind = UNKNOWN
			m.name = parts[0]
		}
		modules[m.name] = m
		for _, n := range strings.Split(parts[1], ", ") {
			if _, found := modules[n]; !found {
				modules[n] = &Module{name: n, kind: UNKNOWN}
			}
		}
	}
	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		sender := modules[strings.TrimLeft(parts[0], "%&")]
		receivers := strings.Split(parts[1], ", ")
		for _, receiver := range receivers {
			modules[receiver].receivesFrom = append(modules[receiver].receivesFrom, sender)
			sender.sendsTo = append(sender.sendsTo, modules[receiver])
		}
	}
	for _, m := range modules {
		if m.kind == CONJUNCTION {
			m.conjState = make(map[string]bool)
			for _, receiver := range m.receivesFrom {
				m.conjState[receiver.name] = OFF
			}
		}
	}
	return modules
}

func allInputHigh(module *Module) bool {
	for _, inSig := range module.conjState {
		if inSig == LOW {
			return false
		}
	}
	return true
}
func solve(modules map[string]*Module, iterations int) int {
	counts := make(map[bool]int)
	for i := 0; i < iterations; i++ {
		signalQueue := []Signal{{from: modules[BUTTON], to: modules[BROADCASTER], kind: LOW}}
		var currSig Signal
		for len(signalQueue) > 0 {
			currSig, signalQueue = signalQueue[0], signalQueue[1:]
			counts[currSig.kind]++
			switch currSig.to.kind {
			case BROADCASTER:
				for _, receiver := range currSig.to.sendsTo {
					signalQueue = append(signalQueue, Signal{from: currSig.to, to: receiver, kind: currSig.kind})
				}
			case CONJUNCTION:
				currSig.to.conjState[currSig.from.name] = currSig.kind
				val := !allInputHigh(currSig.to)
				for _, receiver := range currSig.to.sendsTo {
					signalQueue = append(signalQueue, Signal{from: currSig.to, to: receiver, kind: val})
				}
			case FLIP_FLOP:
				if currSig.kind == LOW {
					currSig.to.flipFlopState = !currSig.to.flipFlopState
					for _, receiver := range currSig.to.sendsTo {
						signalQueue = append(signalQueue, Signal{from: currSig.to, to: receiver, kind: currSig.to.flipFlopState})
					}
				}
			}
		}
	}
	return counts[LOW] * counts[HIGH]
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	lines := strings.Split(data, "\n")
	modules := parseModules(lines)
	fmt.Println(solve(modules, 1000))
}
