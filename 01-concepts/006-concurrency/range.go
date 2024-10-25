package main

import (
	"fmt"
	"time"
)

func rangingChannels() {
	attacks := ""
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	evilNinjas := []string{"Tommy", "Johnny", "Bobby", "Molly"}

	smokeSignal := make(chan string)

	for _, evilNinja := range evilNinjas {
		go attack(evilNinja, smokeSignal)
	}

	for i := 0; i < len(evilNinjas); i++ {
		fmt.Println("Thread finished", <-smokeSignal)
	}
	fmt.Println("My job is done here")
	fmt.Println(attacks)
}

func attack(ninja string, finished chan string) {
	time.Sleep(1 * time.Second)
	fmt.Println("Throwing ninja stars at", ninja)
	finished <- ninja + " killed."
}
