package main

import (
	"fmt"
)

func SayHello(name string) string {
	greetings := "Hello, " + name
	return greetings
}

func main() {
	fmt.Println(SayHello("Suraj"))
}
