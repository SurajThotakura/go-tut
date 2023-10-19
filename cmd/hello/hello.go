package main

import (
	"fmt"
)

const (
	english          = "English"
	ENGLISH_GREETING = "Hello, "

	telugu          = "Telugu"
	TELUGU_GREETING = "Namaskaram, "

	hindi          = "Hindi"
	HINDI_GREETING = "Namaste, "
)

func GetGreeting(language string) string {

	const ENGLISH_GREETING = "Hello, "
	const TELUGU_GREETING = "Namaskaram, "
	const HINDI_GREETING = "Namaste, "

	switch language {
	case english:
		return ENGLISH_GREETING
	case telugu:
		return TELUGU_GREETING
	case hindi:
		return HINDI_GREETING
	default:
		return ENGLISH_GREETING
	}
}

func SayHello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	greeting := GetGreeting(language)
	greetUserString := greeting + name
	return greetUserString
}

func main() {
	fmt.Println(SayHello("Suraj", "English"))
}
