package calculator // for whitebox testing because we can invoke the internal methods

// package calculator_test --> for black box testing (Only the methods that are visible)

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// arrange
// act
// assert

var testcases = []struct{
	name string
	expected float64
	divisor float64
	expectedError error
}{
	{"division", 2.0, 5.0, nil},
	{"division by negative divisor", -2.0, -5.0, nil},
	{"division by negative divisor", 0.0, 0.0, errors.New("division by zero")},
}

func TestDivide(t *testing.T){

	for _,tc := range(testcases){
		t.Run(tc.name, func(t *testing.T){
			
			assert := assert.New(t)

			expected := tc.expected
	
			got, err := Divide(10.0, tc.divisor)

			
			// one way
			if err != nil{
				if err.Error() != tc.expectedError.Error(){
					t.Errorf("expected error %s, got error %s", err.Error(), tc.expectedError.Error())
				}
			}
		
			if got != expected {
				t.Errorf("expected %.1f, found %.1f", expected, got)
			}

			// second way
			assert.Equal(err, tc.expectedError)
			assert.Equal(got, expected)
		})
	}
}

var isZeroTestCases = []struct{
	name string
	expected bool
	arg float64
}{
	{ "argument is zero", true, 0.0 },
	{ "argument is not zero", false, -1.0 },

}
func TestIsZero(t *testing.T){

	for _, tc := range(isZeroTestCases){
		t.Run(tc.name, func(t * testing.T){
			assert := assert.New(t)

			got := isZero(tc.arg)

			assert.Equal(tc.expected, got)
		})
	}

}