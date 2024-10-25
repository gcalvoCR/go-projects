package client_one

import "github.com/gcalvocr/go-patterns/cmd/creationals/singleton/single"

func IncrementAge() {
	p := single.GetInstance()
	p.IngrementAge()
}
