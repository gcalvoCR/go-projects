package single

import (
	"fmt"
	"sync"
)

var (
	p    *person
	once sync.Once
)

func GetInstance() *person {
	once.Do(func() {
		p = &person{}
	})

	// This is not needed when using Once
	// if p == nil {
	//	p = &person{}
	// }

	return p
}

// Watch out!!! The once.Do gets executed only once on per package
func ShowMeAMessage() {
	once.Do(func() {
		fmt.Println("Gimme a Hurray if you see this!")
	})
	fmt.Println("Done!")
}

type person struct {
	name string
	age  int
	mux  sync.Mutex
}

func (p *person) SetName(n string) {
	p.mux.Lock()
	p.name = n
	p.mux.Unlock()
}

func (p *person) GetName() string {
	p.mux.Lock()
	defer p.mux.Unlock()
	return p.name
}

func (p *person) IngrementAge() {
	p.mux.Lock()
	p.age++
	p.mux.Unlock()
}

func (p *person) GetAge() int {
	p.mux.Lock()
	defer p.mux.Unlock()
	return p.age
}
