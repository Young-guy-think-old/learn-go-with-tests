package main

import "fmt"


func main() {
	fmt.Println(hello("Thien"))
}

func hello(name string) string {
	return "Hello, " + name
}