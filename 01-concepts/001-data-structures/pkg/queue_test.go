package pkg_test

import (
	"fmt"
	"testing"

	"github.com/gcalvoCR/data-structures-go/pkg"
)

func Test_Queue(t *testing.T) {
	queue := pkg.Queue{}

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
}
