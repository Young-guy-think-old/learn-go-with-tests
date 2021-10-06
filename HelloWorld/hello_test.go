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
		input := hello(name, "English")
		output := englishHelloPrefix + name
		assertCorrectMessage(t, input, output);
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func (t *testing.T)  {
		input := hello("", "English")
		output := englishHelloPrefix + "world"
		assertCorrectMessage(t, input, output);
	})

	t.Run("in Spanish", func (t *testing.T)  {
		name := "Thien"
		input := hello(name, "Spanish")
		output := spanishHelloPrefix + name
		assertCorrectMessage(t, input, output);
	})
	
}

func TestGreetingPrefix(t *testing.T) {

	assertCorrectMessage := func (t *testing.T, input string, output string) {
		if input != output {
			t.Helper()
			t.Errorf("got %q want %q", input, output)
		}
	}

	t.Run("Spanish", func (t *testing.T)  {
		input := greetingPrefix(spanish)
		output := spanishHelloPrefix
		assertCorrectMessage(t, input, output);
	})

	t.Run("English", func (t *testing.T)  {
		input := greetingPrefix(eng)
		output := englishHelloPrefix
		assertCorrectMessage(t, input, output);
	})

	t.Run("France", func (t *testing.T)  {
		input := greetingPrefix(fre)
		output := franchHelloPrefix
		assertCorrectMessage(t, input, output);
	})
	
}