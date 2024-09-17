package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const japanese = "日本語"
const englishHelloPrefix = "Hello, "
const spanishHolaPrefix = "Hola, "
const frenchBonjourPrefix = "Bonjour, "
const japaneseKonnichihaPrefix = "こんにちは、"


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
	case japanese:
		prefix = japaneseKonnichihaPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
