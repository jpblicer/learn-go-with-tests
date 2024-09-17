package main

import "fmt"

const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const spanishHolaPrefix = "Hola, "


func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == spanish {
		return spanishHolaPrefix + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
