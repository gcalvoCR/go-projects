package functionality

import (
	"fmt"
	"testing"
)

func TestGenerics(t *testing.T) {
	fmt.Println("adding ints:", addInts([]int{1, 2, 4, 5}))
	fmt.Println("adding floats:", addFloats([]float64{1.2, 2.4, 4.1, 5.4}))
	fmt.Println("adding ints with Generics:", addList([]int{1, 2, 4, 5}))
	fmt.Println("adding floats with Generics:", addList([]float64{1.2, 2.4, 4.1, 5.4}))
	fmt.Println("adding strings with Generics:", addList([]string{"test", "with", "strings"}))
	fmt.Println("adding strings with Constraints:", addList([]int{1, 2, 4, 5}))
}
