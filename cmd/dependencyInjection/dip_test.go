package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestHelloWithDIP(t *testing.T) {
	buffer := bytes.Buffer{}
	HelloWithDIP(&buffer, "Suraj")

	got := buffer.String()
	want := "Hello, Suraj"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
