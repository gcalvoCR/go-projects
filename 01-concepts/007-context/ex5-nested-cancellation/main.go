// cancellation with context timeout
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
	t := 1 * time.Second

	ctx, cancel := context.WithTimeout(parent, t)
	defer cancel()

	for i := 0; i < numworkers; i++ {
		wg.Add(1)
		go worker(ctx, i)
	}

	wg.Wait()

}

// worker does work for a random duration of time
func worker(ctx context.Context, id int) {
	defer wg.Done()
	d := time.Duration(r.Int31n(5000)) * time.Millisecond
	fmt.Printf("worker %v started, going to run for %d\n", id, d)

	t := time.Now().Add(500 * time.Millisecond)
	ctx2, cancel := context.WithDeadline(ctx, t)
	go subworker(ctx2, fmt.Sprintf("%d.1", id))
	defer cancel()

	select {
	case <-time.After(d):
	case <-ctx.Done():
		fmt.Printf("Worker %v cancelled\n", id)
		return
	}
	fmt.Printf("Worker %v completed after %d\n", id, d)
}

// worker does work for a random duration of time
func subworker(ctx context.Context, id string) {
	defer wg.Done()
	d := time.Duration(r.Int31n(2000)) * time.Millisecond
	fmt.Printf("subworker %v started, going to run for %d\n", id, d)

	select {
	case <-time.After(d):
	case <-ctx.Done():
		fmt.Printf("Worker %v cancelled\n", id)
		return
	}
	fmt.Printf("Worker %v completed after %d\n", id, d)
}
