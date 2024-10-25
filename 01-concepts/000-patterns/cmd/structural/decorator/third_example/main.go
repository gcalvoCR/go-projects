package main

import "fmt"

type IPersona interface {
	Correr()
	Dormir()
}

type Persona struct{}

func (p Persona) Correr() {
	fmt.Println("Runnn Forest...")
}

func (p Persona) Dormir() {
	fmt.Println("Sleep pequeno saltamontes...")
}

type Maestro struct {
	persona IPersona
}

func (m Maestro) Correr() {
	fmt.Println("El maestro tambien corre")
	m.persona.Correr()
}

func (m Maestro) Dormir() {
	fmt.Println("El maestro tambien duerme")
	m.persona.Dormir()
}

func main() {
	p := Persona{}
	p.Correr()
	p.Dormir()

	m := Maestro{persona: p}
	m.Correr()
	m.Dormir()
}
