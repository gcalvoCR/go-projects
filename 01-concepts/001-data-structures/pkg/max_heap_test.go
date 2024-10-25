package pkg_test

import (
	"fmt"
	"testing"

	"github.com/gcalvoCR/data-structures-go/pkg"
)

func TestInsertMaxHeap(t *testing.T) {
	h := &pkg.MaxHeap{}
	fmt.Println(h)

	heapBuilder := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}

	for _, v := range heapBuilder {
		h.Insert(v)
		fmt.Println(h)
	}
}

func TestExtractMaxHeap(t *testing.T) {
	h := &pkg.MaxHeap{}
	fmt.Println(h)

	heapBuilder := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}

	for _, v := range heapBuilder {
		h.Insert(v)
		fmt.Println(h)
	}

	for i := 0; i < 5; i++ {
		v := h.Extract()
		fmt.Println(v, h)
	}
}
