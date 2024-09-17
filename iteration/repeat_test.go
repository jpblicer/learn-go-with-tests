package iteration

import "testing"

func TestRepeat(t *testing.T) {

	t.Run("Should repeat character five times", func(t *testing.T){
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertCorrectMessage(t, repeated, expected)
	})
	t.Run("Should repeat character 2 times", func(t *testing.T){
		repeated := Repeat("b", 2)
		expected := "bb"
		assertCorrectMessage(t, repeated, expected)
	})
}

func assertCorrectMessage(t testing.TB, repeated, expected string) {
	t.Helper()
	if repeated != expected {
		t.Errorf("Expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", b.N)
	}
}
