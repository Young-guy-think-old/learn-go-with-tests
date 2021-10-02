package main

import "testing"


func TestHello(t *testing.T) {
	name := "Thien"
	input := hello(name)
	output := "Hello, " + name
	if input != output {
		t.Errorf("got %q want %q", input, output)
	}
}