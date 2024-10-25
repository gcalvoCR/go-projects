package main

import (
	"fmt"
	"sync"
)

func waitingForGroupstoFinish() {
	var beeper sync.WaitGroup
	evilNinjas := []string{"Tommy", "Johnny", "Bobby"}
	beeper.Add(len(evilNinjas))
	for _, name := range evilNinjas {
		go attackHim(name, &beeper)
	}
	beeper.Wait()
	fmt.Println("Mission completed")

}

func attackHim(s string, beeper *sync.WaitGroup) {
	fmt.Println(s, "attacked")
	beeper.Done()
}
