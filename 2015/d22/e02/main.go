package main

import "fmt"

type Character struct {
	name         string
	hp, atk, def int
	mana         int
}

type effect func(p, b *Character)

type Spell struct {
	name     string
	duration int
	cost     int
	onStart  effect
	onTick   effect
	onFinish effect
}

func noOp(_, _ *Character) {}

var magicMissile = Spell{
	"Magic Missile",
	1,
	53,
	func(player, boss *Character) {
		boss.hp -= 4
	},
	noOp,
	noOp,
}

var drain = Spell{
	"Drain",
	1,
	73,
	func(player, boss *Character) {
		boss.hp -= 2
		player.hp += 2
	},
	noOp,
	noOp,
}

var shield = Spell{
	"Shield",
	6,
	113,
	func(player, boss *Character) {
		activeEffects["Shield"] = struct{}{}
		player.def += 7
	},
	noOp,
	func(player, boss *Character) {
		delete(activeEffects, "Shield")
		player.def -= 7
	},
}

var poison = Spell{
	"Poison",
	6,
	173,
	func(player, boss *Character) {
		activeEffects["Poison"] = struct{}{}
	},
	func(player, boss *Character) {
		boss.hp -= 3
	},
	func(player, boss *Character) {
		delete(activeEffects, "Poison")
	},
}

var recharge = Spell{
	"Recharge",
	5,
	229,
	func(player, boss *Character) {
		activeEffects["Recharge"] = struct{}{}
	},
	func(player, boss *Character) {
		player.mana += 101
	},
	func(player, boss *Character) {
		delete(activeEffects, "Recharge")
	},
}

// Pick a random spell that won't kill the player
func selectSpell(player *Character) (Spell, bool) {
	var availableSpells = map[string]Spell{
		"Magic Missile": magicMissile,
		"Drain":         drain,
		"Shield":        shield,
		"Poison":        poison,
		"Recharge":      recharge,
	}
	// Discard spells already active or costing too much
	for k, v := range availableSpells {
		if _, found := activeEffects[k]; found || v.cost >= player.mana {
			delete(availableSpells, k)
		}
	}
	// Go shuffles iterating order on maps, just returns the first match
	for _, v := range availableSpells {
		return v, true
	}
	return Spell{}, false
}

var activeEffects = make(map[string]struct{})

func play(player, boss *Character) (int, bool) {
	var cost int
	var queuedEffects [][]effect
	var turnEffects []effect
	for turn := 0; ; turn++ {
		if turn%2 == 0 {
			player.hp--
			if player.hp <= 0 {
				return -1, false
			}
		}
		// Resolve queued effects for the current turn
		if len(queuedEffects) > 1 {
			turnEffects, queuedEffects = queuedEffects[0], queuedEffects[1:]
		} else if len(queuedEffects) == 1 {
			turnEffects = queuedEffects[0]
			queuedEffects = nil
		} else {
			turnEffects = nil
			queuedEffects = nil
		}
		for _, e := range turnEffects {
			e(player, boss)
		}
		// Check if boss is still alive after effects are resolved
		if boss.hp <= 0 {
			return cost, true
		}
		// Boss turn
		if turn%2 == 1 {
			dmg := boss.atk - player.def
			if dmg < 1 {
				dmg = 1
			}
			// Check if player is still alive after boss' turn
			if player.hp -= dmg; player.hp < 0 {
				return -1, false
			}
		} else { // Player turn
			// Select and cast next spell
			spell, found := selectSpell(player)
			if !found {
				return -1, false
			}
			cost += spell.cost
			player.mana -= spell.cost
			spell.onStart(player, boss)
			// Check if boss is still alive after the onStart effect of the spell
			if boss.hp < 0 {
				return cost, true
			}
			for i := len(queuedEffects); i < spell.duration; i++ {
				queuedEffects = append(queuedEffects, []effect{})
			}
			for i := 0; i < spell.duration; i++ {
				queuedEffects[i] = append(queuedEffects[i], spell.onTick)
				if i == spell.duration-1 {
					queuedEffects[i] = append(queuedEffects[i], spell.onFinish)
				}
			}
		}
	}
}

func solve() int {
	var minCost = 99999999999

	// Play 10000 winning games, hope it finds the optimal scenario by chance
	for wonGames := 0; wonGames < 10000; {
		activeEffects = make(map[string]struct{})
		player := Character{"player", 50, 0, 0, 500}
		boss := Character{"boss", 55, 8, 0, 0}
		if cost, won := play(&player, &boss); won {
			wonGames++
			if cost < minCost {
				minCost = cost
			}
		}
	}

	return minCost
}

func main() {
	fmt.Println(solve())
}
