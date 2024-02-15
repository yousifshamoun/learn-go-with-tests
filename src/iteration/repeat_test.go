package iteration

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeating character without specifying length", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := "aaaaa"
		assertCorrectMessage(t, repeated, expected)
	})
	t.Run("repeat character a set number of times", func (t *testing.T) {
		repeated := Repeat("b", 9)
		expected := "bbbbbbbbb"
		assertCorrectMessage(t, repeated, expected)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}