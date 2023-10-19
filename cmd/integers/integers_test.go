package integers

import (
	"fmt"
	"testing"
)

func TestIntegers(t *testing.T) {
	got := Adder(2, 3)
	want := 5

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
func ExampleAdder() {
	sum := Adder(2, 2)
	fmt.Println(sum)
	// Output: 4
	// Not having the colon (Output: 4) broke the code,
	// The Example function was showing up as a separate test and not as a doc example.
}
