package main

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"

	englishGreeting = "Hello, "
	spanishGreeting = "Hola, "
	frenchGreeting  = "Bonjour, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "French":
		prefix = frenchGreeting
	case "Spanish":
		prefix = spanishGreeting
	default:
		prefix = englishGreeting
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
