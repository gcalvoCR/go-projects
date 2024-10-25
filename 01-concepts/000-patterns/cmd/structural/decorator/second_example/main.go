package main

import third_party "github.com/gcalvocr/go-patterns/cmd/decorator/second_example/third-party"

func main() {
	BasicSoldier := third_party.BasicSoldier{}
	third_party.DisplayStats(BasicSoldier)

	soldierWithSword := SoldierWithSword{soldier: BasicSoldier}
	third_party.DisplayStats(soldierWithSword)
}

// Decorator 1: Soldier with a sword

type SoldierWithSword struct {
	soldier third_party.ISoldier
}

func (s SoldierWithSword) Attack() int {
	return s.soldier.Attack() + 10
}

func (s SoldierWithSword) Defense() int {
	return s.soldier.Defense()
}
