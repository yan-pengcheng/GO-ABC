package main

import (
	"fmt"
)

const spanish = "Spanish"
const japanese = "japanese"
const helloPrefix = "Hello, "
const japaneseHelloPrefix = "ようこそ, "
const englishHelloPrefix = "Hello, "

// Hello .
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == japanese {
		return japaneseHelloPrefix + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
