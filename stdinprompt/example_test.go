package stdinprompt_test

import (
	"bytes"
	"fmt"
	"os"

	"github.com/lwithers/pkg/stdinprompt"
)

func ExampleNew() {
	q := make([]byte, 4096)
	in := stdinprompt.New()
	for {
		n, err := in.Read(q)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		os.Stdout.Write(bytes.ToUpper(q[:n]))
	}
}
