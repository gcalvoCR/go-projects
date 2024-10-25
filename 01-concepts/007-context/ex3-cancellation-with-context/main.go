// cancellation with context cancel
package main

import (
	"context"
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
	parent := context.Background()

	ctx, cancel := context.WithCancel(parent)

	for i := 0; i < numworkers; i++ {
		wg.Add(1)
		go worker(ctx, i)
	}
	<-time.After(1 * time.Second)
	cancel() // this is the way we send the context info via the channel

	wg.Wait()

}

// worker does work for a random duration of time
func worker(ctx context.Context, id int) {
	defer wg.Done()
	d := time.Duration(r.Int31n(5000)) * time.Millisecond
	fmt.Printf("worker %v started, going to run for %d\n", id, d)

	select {
	case <-time.After(d):
	case <-ctx.Done():
		fmt.Printf("Worker %v cancelled\n", id)
		return
	}
	fmt.Printf("Worker %v completed after %d\n", id, d)
}
