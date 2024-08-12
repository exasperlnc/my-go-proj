package main

import "fmt"

func main() {
	fmt.Println(Hello("Logan"))
}

func Hello(name string) string {
	englishHelloPrefix := "Hello, "
	exclamation := "!"
	if name == "" {
		name = "World"
	} 
	return englishHelloPrefix + name + exclamation
}