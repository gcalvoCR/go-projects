package main

import (
	"fmt"
	"time"
)

func simple() {
	now := time.Now()
	defer func() {
		fmt.Println("Elapsed time", time.Since(now))
	}()

	name := make(chan string)
	go func() {
		name <- "Hello world"
	}()

	fmt.Println(<-name)

}
