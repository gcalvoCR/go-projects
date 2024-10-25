package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	s  = rand.NewSource(time.Now().Unix())
	r  = rand.New(s)
	wg = &sync.WaitGroup{}
)

const numworkers = 3

func main() {

	for i := 0; i < numworkers; i++ {
		wg.Add(1)
		go worker(i)
	}
	wg.Wait()

}

// worker does work for a random duration of time
func worker(id int) {
	d := time.Duration(r.Int31n(5000)) * time.Millisecond
	fmt.Printf("worker %v started, going to run for %d\n", id, d)
	time.Sleep(d)
	fmt.Printf("worker %v completed after %d\n", id, d)
	wg.Done()
}
