package main

import "fmt"

func main() {
	fmt.Println(Hello("Logan"))
}

func Hello(name string) string {
	return "Hello, " + name + "!"
}