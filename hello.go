package main

import "fmt"

func main() {
	fmt.Println(Hello("Logan", ""))
}
const exclamation = "!"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const spanish = "Spanish"

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	} 
	if language == spanish {
		return spanishHelloPrefix + name + exclamation
	}
	return englishHelloPrefix + name + exclamation
}