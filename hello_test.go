package main

import "testing"


func TestHello(t *testing.T) {

	assertCorrectMessage := func (t *testing.T, input string, output string) {
		if input != output {
			t.Helper()
			t.Errorf("got %q want %q", input, output)
		}
	}

	t.Run("saying hello to people", func (t *testing.T)  {
		name := "Thien"
		input := hello(name)
		output := englishHelloPrefix + name
		assertCorrectMessage(t, input, output);
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func (t *testing.T)  {
		input := hello("")
		output := englishHelloPrefix + "world"
		assertCorrectMessage(t, input, output);
	})
}