package main

import "fmt"


var pointer *string

func main() {
	
	myName:= "Gabriel"

	pointer = &myName
	
	fmt.Println(myName)
	changeMyName(myName)
	fmt.Println(myName)
	changeMyNameWithPointer(&myName) // passes the address
	fmt.Println(myName)

	fmt.Println(pointer)  // pointer is the position in memory
	fmt.Println(&myName)  // sending the position in memory
	fmt.Println(*pointer) // dereferecing
}

func changeMyName(name string){
	name="The warrior"
}

func changeMyNameWithPointer(name *string){
	*name="The warrior"
}