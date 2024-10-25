package pkg_test

import (
	"fmt"
	"testing"

	"github.com/gcalvoCR/data-structures-go/pkg"
)

func Test_(t *testing.T) {
	stack := pkg.Stack{}

	// Pushing elements to the stack
	stack.Push(2)
	stack.Push(4)
	stack.Push(6)

	// Popping elements from the stack
	fmt.Println(stack.Pop()) // Outputs: 6
	fmt.Println(stack.Pop()) // Outputs: 4
	fmt.Println(stack.Pop()) // Outputs: 2
	fmt.Println(stack.Pop()) // Outputs: -1

}
