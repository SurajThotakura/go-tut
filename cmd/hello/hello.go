package main

import (
	"fmt"
)

func SayHello(name string) string {
	if name == "" {
		name = "world"
	}
	const greeting = "Hello, "
	greetUserString := greeting + name
	return greetUserString
}

func main() {
	fmt.Println(SayHello("Suraj"))
}
