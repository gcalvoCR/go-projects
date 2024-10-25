package third_party

import "log"

type ISoldier interface {
	Attack() int
	Defense() int
}

type BasicSoldier struct{}

func (b BasicSoldier) Attack() int {
	return 1
}

func (b BasicSoldier) Defense() int {
	return 1
}

func DisplayStats(soldier ISoldier) {
	log.Printf("Soldier stats: attack %d, defense %d",
		soldier.Attack(),
		soldier.Defense())
}
