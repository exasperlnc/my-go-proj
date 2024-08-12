package main

import "fmt"

func main() {
	fmt.Println(Hello("Logan"))
}

func Hello(name string) string {
	englishHelloPrefix := "Hello, "
	exclamation := "!"
	return englishHelloPrefix + name + exclamation
}