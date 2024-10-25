package main

import "fmt"

type LogRegistry struct {
	Handler Decorator
}

func NewLogRegistry(handler Decorator) *LogRegistry {
	return &LogRegistry{Handler: handler}
}

func (lr *LogRegistry) Process() error {
	fmt.Println("peticion")
	return lr.Handler.Process()
}
