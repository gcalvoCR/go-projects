package main

import (
	"fmt"
	"os"
)

func main() {
	route := NewRoute()
	start(&route)

	var path string
	_, err := fmt.Scanf("%s", &path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	route.Exec(path)
}

func start(route *Route) {
	route.Add(NewLogRegistry(NewHandlerSayHello()), "/saludar")
	route.Add(NewLogRegistry(NewHandlerSayGoodBye()), "/despedirse")
}
