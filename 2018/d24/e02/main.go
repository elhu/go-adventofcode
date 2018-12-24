package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Army struct {
	groups []*Group
}

func (a *Army) cleanup() {
	res := make([]*Group, 0)
	for _, g := range a.groups {
		if g.target != nil {
			g.target.selected = false
		}
		g.target = nil
		if g.unitCount > 0 {
			res = append(res, g)
		}
	}
	a.groups = res
}

func (a *Army) unitCount() int {
	res := 0
	for _, g := range a.groups {
		res += g.unitCount
	}
	return res
}

const (
	ImmuneSystem = iota
	Infection    = iota
)

const (
	Bludgeoning = iota
	Slashing    = iota
	Cold        = iota
	Radiation   = iota
	Fire        = iota
)

type sortGroupTargetSelection []*Group

func (s sortGroupTargetSelection) Less(i, j int) bool {
	if s[i].effectivePower() != s[j].effectivePower() {
		return s[i].effectivePower() < s[j].effectivePower()
	}
	return s[i].initiative < s[j].initiative
}

func (s sortGroupTargetSelection) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortGroupTargetSelection) Len() int {
	return len(s)
}

type sortGroupTargetAttack []*Group

func (s sortGroupTargetAttack) Less(i, j int) bool {
	return s[i].initiative < s[j].initiative
}

func (s sortGroupTargetAttack) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortGroupTargetAttack) Len() int {
	return len(s)
}

type Group struct {
	id           int
	army         byte
	unitCount    int
	unitHitPoint int
	weakTo       []int
	immuneTo     []int
	initiative   int
	damageType   int
	unitDamage   int
	target       *Group
	selected     bool
}

func (g *Group) isImmuneTo(elem int) bool {
	for _, e := range g.immuneTo {
		if e == elem {
			return true
		}
	}
	return false
}

func (g *Group) isWeakTo(elem int) bool {
	for _, e := range g.weakTo {
		if e == elem {
			return true
		}
	}
	return false
}

func (g *Group) effectivePower() int {
	b := 0
	if g.army == ImmuneSystem {
		b = boost
	}
	return g.unitCount * (g.unitDamage + b)
}

func (g *Group) damageScore(o *Group) int {
	if o.isImmuneTo(g.damageType) {
		return 0
	} else if o.isWeakTo(g.damageType) {
		return g.effectivePower() * 2
	}
	return g.effectivePower()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (g *Group) attack() {
	if g.target != nil {
		unitsKilled := min(g.damageScore(g.target)/g.target.unitHitPoint, g.target.unitCount)
		g.target.unitCount -= unitsKilled
	}
}

var unitExp = regexp.MustCompile(`(\d+) units each with (\d+) hit points.+with an attack that does (\d+) (\w+) damage at initiative (\d+)`)
var specialExp = regexp.MustCompile(`\((.+)\)`)

var matchType = map[string]int{
	"bludgeoning": Bludgeoning,
	"slashing":    Slashing,
	"cold":        Cold,
	"radiation":   Radiation,
	"fire":        Fire,
}

var printableType = map[int]string{
	Bludgeoning: "bludgeoning",
	Slashing:    "slashing",
	Cold:        "cold",
	Radiation:   "radiation",
	Fire:        "fire",
}

var printableArmy = map[byte]string{
	Infection:    "Infection",
	ImmuneSystem: "Immune system",
}

func parseArmies(input [][]byte) (*Army, *Army) {
	immuneSystem := &Army{groups: make([]*Group, 0)}
	infection := &Army{groups: make([]*Group, 0)}
	var currentArmy *Army
	var armyID byte
	var id int
	for _, l := range input {
		line := string(l)
		if line == "Immune System:" {
			currentArmy = immuneSystem
			armyID = ImmuneSystem
			id = 1
		} else if line == "Infection:" {
			currentArmy = infection
			armyID = Infection
			id = 1
		}
		if match := unitExp.FindStringSubmatch(line); match != nil {
			unitCount, _ := strconv.Atoi(match[1])
			unitHitPoint, _ := strconv.Atoi(match[2])
			unitDamage, _ := strconv.Atoi(match[3])
			initiative, _ := strconv.Atoi(match[5])
			group := &Group{
				id:           id,
				army:         armyID,
				unitCount:    unitCount,
				unitHitPoint: unitHitPoint,
				unitDamage:   unitDamage,
				damageType:   matchType[match[4]],
				initiative:   initiative,
			}
			id++
			if specialMatch := specialExp.FindStringSubmatch(line); specialMatch != nil {
				special := specialMatch[1]
				for _, s := range strings.Split(special, "; ") {
					words := strings.Split(s, " ")
					types := make([]int, 0)
					for _, w := range words[2:] {
						t, exists := matchType[strings.TrimRight(w, ",")]
						if !exists {
							panic("WTF")
						}
						types = append(types, t)
					}
					if words[0] == "immune" {
						group.immuneTo = types
					} else if words[0] == "weak" {
						group.weakTo = types
					} else {
						panic(words[0])
					}
				}
			}
			currentArmy.groups = append(currentArmy.groups, group)
		}
	}
	return immuneSystem, infection
}

func selectTargets(immuneSystem, infection *Army) {
	orderedGroups := append(make([]*Group, 0), immuneSystem.groups...)
	orderedGroups = append(orderedGroups, infection.groups...)
	sort.Sort(sort.Reverse(sortGroupTargetSelection(orderedGroups)))

	for _, g := range orderedGroups {
		maxDamage := -1
		maxInitiative := -1
		maxEP := -1
		var target *Group
		var targetArmy *Army
		if g.army == ImmuneSystem {
			targetArmy = infection
		} else {
			targetArmy = immuneSystem
		}
		for _, t := range targetArmy.groups {
			ds := g.damageScore(t)
			if !t.selected {
				if ds > 0 && ds > maxDamage || (ds == maxDamage && t.effectivePower() > maxEP) || (ds == maxDamage && t.effectivePower() == maxEP && t.initiative > maxInitiative) {
					maxDamage = ds
					maxEP = t.effectivePower()
					maxInitiative = t.initiative
					target = t
				}
			}
		}
		if maxDamage > 0 {
			target.selected = true
			g.target = target
		}
	}
}

func attack(immuneSystem, infection *Army) {
	orderedGroups := append(make([]*Group, 0), immuneSystem.groups...)
	orderedGroups = append(orderedGroups, infection.groups...)
	sort.Sort(sort.Reverse(sortGroupTargetAttack(orderedGroups)))
	for _, g := range orderedGroups {
		g.attack()
	}
}

func solve(immuneSystem, infection *Army) int {
	res := 0
	for _, g := range append(immuneSystem.groups, infection.groups...) {
		res += g.unitCount
	}
	return res
}

var boost = 0

func play(immuneSystem, infection *Army) bool {
	round := 0
	for len(immuneSystem.groups) > 0 && len(infection.groups) > 0 {
		round++
		prevCount := immuneSystem.unitCount() + infection.unitCount()
		selectTargets(immuneSystem, infection)
		attack(immuneSystem, infection)
		immuneSystem.cleanup()
		infection.cleanup()
		if immuneSystem.unitCount()+infection.unitCount() == prevCount {
			return false
		}
	}
	return true
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	data = bytes.TrimSuffix(data, []byte{'\n'})
	lines := bytes.Split(data, []byte{'\n'})
	for {
		boost += 1
		immuneSystem, infection := parseArmies(lines)
		completed := play(immuneSystem, infection)
		if completed && len(immuneSystem.groups) > 0 {
			fmt.Println(solve(immuneSystem, infection))
			break
		}
	}
}
