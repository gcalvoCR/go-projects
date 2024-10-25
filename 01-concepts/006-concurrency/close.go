package main

import (
	"fmt"
	"math/rand"
	"time"
)

func closinChannels() {
	now := time.Now()
	defer func() {
		fmt.Println("Elapsed time", time.Since(now))
	}()

	channel := make(chan string)
	go throwringNinjaStar(channel)

	for message := range channel {
		fmt.Println(message)
	}

	for {
		message, open := <-channel
		if !open {
			break
		}
		fmt.Println(message)
	}

}

func throwringNinjaStar(channel chan string) {
	rounds := 3
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < rounds; i++ {
		score := rand.Intn(10)
		time.Sleep(1 * time.Second)
		channel <- fmt.Sprintln("You scored ", score)
	}
	close(channel)
}
