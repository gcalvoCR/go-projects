package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	lock  sync.Mutex
	count = 0
)

func usingMutex() {
	for i := 0; i < 1000; i++ {
		go increment()
	}
	time.Sleep(2 * time.Second)
	fmt.Println(count)
}

func increment() {
	lock.Lock()
	count++
	lock.Unlock()
}
