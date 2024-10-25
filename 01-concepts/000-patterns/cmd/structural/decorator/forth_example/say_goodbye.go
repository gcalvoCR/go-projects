package main

import "fmt"

type HandlerSayGoodBye struct{}

func NewHandlerSayGoodBye() *HandlerSayGoodBye {
	return &HandlerSayGoodBye{}
}

func (h HandlerSayGoodBye) Process() error {
	fmt.Println("Adios")
	return nil
}
