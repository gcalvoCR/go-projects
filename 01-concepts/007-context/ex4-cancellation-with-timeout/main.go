// cancellation with context timeout
package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

var (
	s = rand.NewSource(time.Now().Unix())
	r = rand.New(s)
)

const numworkers = 3

func main() {
	parent := context.Background()
	t := 1 * time.Second

	ctx, cancel := context.WithTimeout(parent, t)
	defer cancel()

	for i := 0; i < numworkers; i++ {
		go worker(ctx, i)
	}

	<-ctx.Done()

}

// worker does work for a random duration of time
func worker(ctx context.Context, id int) {
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
