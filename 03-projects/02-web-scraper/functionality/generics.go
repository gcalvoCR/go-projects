package functionality

import (
	"golang.org/x/exp/constraints"
)

// Use case
// When we need to write the same function for diffent types

func addInts(list []int) int {
	var sum int

	for _, v := range list {
		sum += v
	}

	return sum
}

func addFloats(list []float64) float64 {
	var sum float64

	for _, v := range list {
		sum += v
	}

	return sum
}

type NumberType interface {
	int | float64 | float32 | string
}

func addList[T NumberType](list []T) T {
	var sum T

	for _, v := range list {
		sum += v
	}

	return sum
}

func addListWithConstraints[T constraints.Integer](list []T) T {
	var sum T

	for _, v := range list {
		sum += v
	}

	return sum
}
