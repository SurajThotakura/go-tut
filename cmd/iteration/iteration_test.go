package iteration

import (
	"fmt"
	"testing"
)

func TestIteration(t *testing.T) {
	got := Iteration("p", 4)
	want := "pppp"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
func ExampleIteration() {
	iteration := Iteration("o", 3)
	fmt.Println(iteration)
	// Output: ooo
}
