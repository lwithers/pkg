/*
Package stdinprompt provides an io.Reader that will prompt the user if nothing
is read within a short period of time. It is useful for programs which default
to reading from stdin if no commandline arguments are provided, as it can
indicate to the user that a program is not actually performing any action until
data arrives. The prompt is not displayed if data becomes available before the
timeout.
*/
package stdinprompt

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	// DefaultPromptTime is the timeout period after which, if no data has
	// been read, we will display a prompt.
	DefaultPromptTime = 250 * time.Millisecond

	// StdinPromptMsg is the default message displayed.
	StdinPromptMsg = "Waiting for data on stdin."
)

// New returns a new prompting reader for stdin. The prompt will be written to
// stderr.
func New() io.Reader {
	return NewEx(os.Stdin, DefaultPromptTime, os.Stderr, StdinPromptMsg)
}

// NewEx returns a new prompting reader. The source may be specified along with
// the prompt message, timeout and destination.
func NewEx(raw io.Reader, when time.Duration, term io.Writer, msg string,
) io.Reader {
	pr := &prompter{
		raw: raw,
		tmr: time.AfterFunc(when, func() {
			fmt.Fprintln(term, msg)
		}),
	}
	return pr
}

type prompter struct {
	raw io.Reader
	tmr *time.Timer
}

func (pr *prompter) Read(buf []byte) (int, error) {
	n, err := pr.raw.Read(buf)
	if n > 0 {
		pr.tmr.Stop()
	}
	return n, err
}
