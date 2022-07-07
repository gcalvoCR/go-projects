package main

import (
	"fmt"
	"net/url"
	"strings"
)



func firstFunc(){
	form := url.Values{}

	form.Add("apply", "true")
	form.Add("username", "admin")
	form.Add("password", "hello")
	form.Add("propertylist", "username,password")

	fmt.Println(form)

	fmt.Println(strings.NewReader(form.Encode()))

}

type InstanceType string

const (
	Author InstanceType = "author"
	Publish InstanceType = "publish"
)

func main(){
	print(Author)
}

func print(instanceType InstanceType){
	fmt.Println(instanceType)
}