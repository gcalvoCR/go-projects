package main

import "fmt"

type Username string

type Age int

type person struct {
	username Username
	email    string
}

func main() {
	p := person{
		username: Username("Gabriel"),
		email:    "hello",
	}
	age := Age(20)
	fmt.Println(p)
	fmt.Println(age == 20)
}
