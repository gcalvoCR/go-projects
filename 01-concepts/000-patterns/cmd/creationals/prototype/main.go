package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name          string
	Age           int
	IsMarried     bool
	FavoriteColor *string
	PhoneNumbers  []string
	Attributes    map[string]string
}

func (p *Person) Clone() Person {
	phoneNumbers := append([]string{}, p.PhoneNumbers...)
	favouriteColor := *p.FavoriteColor
	attributes := map[string]string{}

	for k, v := range p.Attributes {
		attributes[k] = v
	}

	return Person{
		Name:          p.Name,
		Age:           p.Age,
		IsMarried:     p.IsMarried,
		FavoriteColor: &favouriteColor,
		PhoneNumbers:  phoneNumbers,
		Attributes:    attributes,
	}
}

func main() {

	orange := "Orange"
	red := "Red"

	gabriel := Person{
		Name:          "Gabriel",
		Age:           36,
		IsMarried:     true,
		FavoriteColor: &orange,
		PhoneNumbers:  []string{"62805172"},
		Attributes: map[string]string{
			"degree":    "master",
			"character": "superman",
		},
	}

	pablo := gabriel.Clone()

	pablo.Name = "Pablo"
	pablo.Age = 35
	pablo.FavoriteColor = &red
	pablo.PhoneNumbers[0] = "86179853"
	pablo.Attributes["character"] = "spiderman"

	jsonDataGabriel, _ := json.MarshalIndent(&gabriel, "", "	")
	fmt.Println(string(jsonDataGabriel))

	jsonDataPablo, _ := json.MarshalIndent(&pablo, "", "	")
	fmt.Println(string(jsonDataPablo))
}
