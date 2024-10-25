package pkg

import "fmt"

// Stack structure using a slice
type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
	fmt.Println("element added:", item)
}

func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		fmt.Println("The stack is empty")
		return -1 // Error: Stack is empty
	}

	lastIndex := len(s.items) - 1
	popped := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return popped
}
