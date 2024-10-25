package main

import (
	"fmt"
	"sync"

	"github.com/gcalvocr/go-patterns/cmd/creationals/singleton/client_one"
	"github.com/gcalvocr/go-patterns/cmd/creationals/singleton/client_two"
	"github.com/gcalvocr/go-patterns/cmd/creationals/singleton/single"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(200)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			client_one.IncrementAge()
		}()
		go func() {
			defer wg.Done()
			client_two.IncrementAge()
		}()
	}
	wg.Wait()
	p := single.GetInstance()
	fmt.Println("edad", p.GetAge())

	// Watch out!!! The once.Do gets executed only once on per package
	single.ShowMeAMessage()
}
