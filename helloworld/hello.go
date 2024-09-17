package main

import "fmt"

const (
	spanish = "Spanish"
	french = "French"
	japanese = "日本語"

	englishHelloPrefix = "Hello, "
	spanishHolaPrefix = "Hola, "
	frenchBonjourPrefix = "Bonjour, "
	japaneseKonnichihaPrefix = "こんにちは、"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHolaPrefix
	case french:
		prefix = frenchBonjourPrefix
	case japanese:
		prefix = japaneseKonnichihaPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
