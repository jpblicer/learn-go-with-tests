package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func (t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when empty string is is supplied", func (t *testing.T) {
		got := Hello("", "")
		want:= "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Sylvain", "French")
		want := "Bonjour, Sylvain"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Japanese", func(t *testing.T){
		got:= Hello("ももこ", "日本語")
		want := "こんにちは、ももこ"
		assertCorrectMessage(t, got, want)
	})

}


func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
