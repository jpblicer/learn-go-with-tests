package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}

	t.Run("known word", func(t *testing.T){
		got := dictionary.Search("test")
		want := "this is a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func (t *testing.T)  {
		_, err := dictionary.Search("unknown")
		want := "could not find word you are looking for"

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertStrings(t, err.Error(), want)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
