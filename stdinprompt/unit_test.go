package stdinprompt

import (
	"bytes"
	"strings"
	"testing"
	"time"
)

func TestNoPrompt(t *testing.T) {
	// set up a buffer (which supports io.Reader) with some test data
	exp := []byte{0xa, 0xb, 0xc, 0xd}
	in := bytes.NewBuffer(nil)
	in.Write(exp)

	// set up a message capture buffer
	msgcap := bytes.NewBuffer(nil)

	// set up the prompter
	inpr := NewEx(in, DefaultPromptTime, msgcap, StdinPromptMsg)

	// read test data
	out := make([]byte, len(exp))
	n, err := inpr.Read(out)
	if n != len(exp) {
		t.Errorf("read %d bytes, expected %d", n, len(exp))
	}
	if err != nil {
		t.Errorf("unexpected read error: %v", err)
	}
	if !bytes.Equal(out, exp) {
		t.Errorf("read mismatch: got %X expected %X", out, exp)
	}

	// wait for at least the prompt timer
	time.Sleep(2 * DefaultPromptTime)

	// verify we didn't get anything in our capture buffer
	if msgcap.Len() > 0 {
		t.Errorf("got unexpected error %q", msgcap.String())
	}
}

func TestPrompt(t *testing.T) {
	// set up an (unused) input buffer
	in := bytes.NewBuffer(nil)

	// set up a message capture buffer
	msgcap := bytes.NewBuffer(nil)

	// set up the prompter
	_ = NewEx(in, DefaultPromptTime, msgcap, StdinPromptMsg)

	// wait for at least the prompt timer
	time.Sleep(2 * DefaultPromptTime)

	// verify that we received the expected message in the capture buffer
	switch {
	case msgcap.Len() == 0:
		t.Errorf("no prompt was received")
	case strings.TrimSpace(msgcap.String()) != StdinPromptMsg:
		t.Errorf("got unexpected message %q", msgcap.String())
	}
}
