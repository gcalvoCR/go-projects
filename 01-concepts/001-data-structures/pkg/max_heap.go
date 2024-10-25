package pkg

import "fmt"

/*
 A heap is a data struct that be expressed as a complete tree
 a complete tree means that all levels are full except for the lowest level
 where some nodes can be empty

 Max heaps, the larget key will be stored in the root

 Max heaps are useful when we want to remove the highest key (priority queues)

 Usually under the hood the keys are stored in arrays because
 it is possible to easily calculate the indeces of the children, based on the parents index

 parent_index * 2 + 1 = left_child_index
 parent_index * 2 + 2 = right_child_index

 Insert process

 we insert at the bottom to the right of the tree, getting the last index
 then we rearrange the nodes
 we compare the new node and swap it if necessary

 the process of rearranging the nodes are heapify

 To extract a key
 we usually just extract the root
 we fill the root with the last node
 and then rearrange

 Time complexity
 O(h)
 O(log n)

*/

// Max heap struct has a slice that holds the array
type MaxHeap struct {
	array []int
}

// Insert adds an element to the heap
func (h *MaxHeap) Insert(key int) {
	// add an element
	h.array = append(h.array, key)
	// heapify
	h.maxHeapifyUp(len(h.array) - 1)
}

func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	// handle struct emtpy
	if len(h.array) == 0 {
		fmt.Println("Cannot extract because structure is empty")
		return -1
	}

	// take out the last index and put it in the root
	lastIndex := len(h.array) - 1
	h.array[0] = h.array[lastIndex]
	h.array = h.array[:lastIndex]

	// maxHeeapifyDown will heapify from the top down
	h.maxHeapifyDown(0)

	return extracted
}

// maxHeapifyUp will heapify from the bottom to the top
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		// Swapping
		h.swap(parent(index), index)
		// updating index
		index = parent(index)
	}
}

// maxHeapifyDown will heapify from the bottom to the top
func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	// loop while index has at least one child

	for l <= lastIndex {
		if l == lastIndex { // when left child is the only child
			childToCompare = l
		} else if h.array[l] > h.array[r] { // when left child is larger
			childToCompare = l
		} else { // when right child is larger
			childToCompare = r
		}

		// compare array value of current index to larger child and swap if smaller
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}

	}
}

// get the parent index
func parent(i int) int {
	return (i - 1) / 2
}

// get the left child index
func left(i int) int {
	return i*2 + 1
}

// get the right child index
func right(i int) int {
	return i*2 + 2
}

func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

// Extract returns the largets key and removes it from the heap
