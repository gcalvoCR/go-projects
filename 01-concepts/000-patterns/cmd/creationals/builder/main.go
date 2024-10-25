package main

import "fmt"

// Car struct represents the object we're building
type Car struct {
	Make  string
	Model string
	Year  int
	Color string
}

// CarBuilder struct holds the properties needed to build a Car
type CarBuilder struct {
	make  string
	model string
	year  int
	color string
}

// NewCarBuilder initializes a new CarBuilder with default values
func NewCarBuilder() *CarBuilder {
	return &CarBuilder{}
}

// SetMake sets the make of the car
func (cb *CarBuilder) SetMake(make string) *CarBuilder {
	cb.make = make
	return cb
}

// SetModel sets the model of the car
func (cb *CarBuilder) SetModel(model string) *CarBuilder {
	cb.model = model
	return cb
}

// SetYear sets the year of the car
func (cb *CarBuilder) SetYear(year int) *CarBuilder {
	cb.year = year
	return cb
}

// SetColor sets the color of the car
func (cb *CarBuilder) SetColor(color string) *CarBuilder {
	cb.color = color
	return cb
}

// Build constructs and returns the final Car object
func (cb *CarBuilder) Build() Car {
	return Car{
		Make:  cb.make,
		Model: cb.model,
		Year:  cb.year,
		Color: cb.color,
	}
}

func main() {
	// Create a car using the builder pattern
	car := NewCarBuilder().
		SetMake("Tesla").
		SetModel("Model S").
		SetYear(2023).
		SetColor("Red").
		Build()

	fmt.Printf("Car: %+v\n", car)
}
