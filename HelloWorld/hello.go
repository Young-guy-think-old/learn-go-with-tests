package main

import "fmt"


const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const franchHelloPrefix = "Bonjour, "

const eng = "English"
const spanish = "Spanish"
const fre = "France"



func main() {
	fmt.Println(hello("Thien", "Spanish"))
}

func hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	prefix := greetingPrefix(language);

	return prefix + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case eng:
		prefix = englishHelloPrefix
	case fre:
		prefix = franchHelloPrefix
	}

	return;
}