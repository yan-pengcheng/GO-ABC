package main

import (
	"fmt"
)

const englishHelloPrefix = "Hello, "

//Hello print
func Hello(name string) string {
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
