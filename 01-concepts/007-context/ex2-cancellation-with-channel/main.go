// cancellation with channels
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
	ch := make(chan bool, numworkers)

	for i := 0; i < numworkers; i++ {
		wg.Add(1)
		go worker(ch, i)
	}
	// we wait for a second and then send a message on the channel
	<-time.After(1 * time.Second)

	// we send a message to those channels to cancell them
	for i := 0; i < numworkers; i++ {
		ch <- true
	}
	wg.Wait()

}

// worker does work for a random duration of time
func worker(ch chan bool, id int) {
	defer wg.Done()

	d := time.Duration(r.Int31n(5000)) * time.Millisecond
	fmt.Printf("worker %v started, going to run for %d\n", id, d)

	select { // the way you can listen on multiple channels
	case <-time.After(d): // the time package has some channels by default
	case <-ch:
		fmt.Printf("Worker %v cancelled\n", id)
		return
	}
	fmt.Printf("Worker %v completed after %d\n", id, d)
}
