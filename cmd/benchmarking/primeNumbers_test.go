package benchmarking

import "testing"

func TestCheckPrimeNumber(t *testing.T) {
	got := checkPrimeNumber(1)
	want := true

	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}
func TestPrintPrimesTill(t *testing.T) {
	got := printPrimesTill(6)
	want := "1, 2, 3, 5"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
func BenchmarkPrimesTill(b *testing.B) {
	for i := 0; i < b.N; i++ {
		printPrimesTill(1000)
	}
}
