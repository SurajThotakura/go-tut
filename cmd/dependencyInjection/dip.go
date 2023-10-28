package dependencyinjection

import (
	"bytes"
	"fmt"
)

func HelloWithDIP(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
