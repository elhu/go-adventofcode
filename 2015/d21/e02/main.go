// No code change, changed the input instead

package main

import (
	"fmt"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func atoi(str string) int {
	i, e := strconv.Atoi(str)
	check(e)
	return i
}

type Item struct {
	name   string
	cost   int
	damage int
	armor  int
}

var weapons = []Item{
	{"Dagger", 8, 4, 0},
	{"Shortsword", 10, 5, 0},
	{"Warhammer", 25, 6, 0},
	{"Longsword", 40, 7, 0},
	{"Greataxe", 74, 8, 0},
}

var armors = []Item{
	{"Leather", 13, 0, 1},
	{"Chainmail", 31, 0, 2},
	{"Splintmail", 53, 0, 3},
	{"Bandedmail", 75, 0, 4},
	{"Platemail", 102, 0, 5},
}

var rings = []Item{
	{"Damage +1", 25, 1, 0},
	{"Damage +2", 50, 2, 0},
	{"Damage +3", 100, 3, 0},
	{"Defense +1", 20, 0, 1},
	{"Defense +2", 40, 0, 2},
	{"Defense +3", 80, 0, 3},
}

type Character struct {
	name         string
	hp, atk, def int
}

func equipments() [][]Item {
	var res [][]Item

	armorCombo := [][]Item{{}}
	for _, a := range armors {
		armorCombo = append(armorCombo, []Item{a})
	}

	ringsCombo := [][]Item{{}}
	for i, r := range rings {
		ringsCombo = append(ringsCombo, []Item{r})
		for j := i + 1; j < len(rings); j++ {
			ringsCombo = append(ringsCombo, []Item{r, rings[j]})
		}
	}

	for _, w := range weapons {
		for _, a := range armorCombo {
			for _, r := range ringsCombo {
				set := append(append([]Item{w}, a...), r...)
				res = append(res, set)
			}
		}
	}

	return res
}

func play(player, boss *Character) *Character {
	for i := 0; ; i++ {
		attacker, defender := boss, player
		if i%2 == 0 {
			attacker, defender = player, boss
		}
		damage := attacker.atk - defender.def
		if damage < 1 {
			damage = 1
		}
		defender.hp -= damage
		if defender.hp <= 0 {
			return attacker
		}
	}
}

func main() {
	maxCost := 0
	for _, set := range equipments() {
		cost := 0
		def := 0
		atk := 0
		for _, e := range set {
			cost += e.cost
			def += e.armor
			atk += e.damage
		}
		boss := Character{"boss", 100, 8, 2}
		player := Character{"player", 100, atk, def}
		if winner := play(&player, &boss); winner.name == "boss" && cost > maxCost {
			maxCost = cost
		}
	}
	fmt.Println(maxCost)
}
