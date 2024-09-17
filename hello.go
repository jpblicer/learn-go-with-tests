package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHolaPrefix = "Hola, "
const frenchBonjourPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case spanish:
		prefix = spanishHolaPrefix
	case french:
		prefix = frenchBonjourPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
