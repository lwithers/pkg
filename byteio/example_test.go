package byteio_test

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"os"

	"github.com/lwithers/pkg/byteio"
)

// SwapEndian32 swaps the endianness of a stream of uint32 integers.
//
// Note this function takes standard io.Reader and io.Writer interfaces, and
// uses byteio to (possibly) wrap these with bufio.Reader/Writer for efficient
// byte-oriented operation. This is transparent to the caller.
func SwapEndian32(in io.Reader, out io.Writer) error {
	// optionally wrap input and output
	bin, bout := byteio.NewReader(in), byteio.NewWriter(out)

	// since output may be buffered, make sure to flush it when leaving
	// this function
	defer byteio.FlushIfNecessary(bout)

	for {
		// it doesn't matter whether we read BE or LE, as long as we
		// write the opposite!
		n, err := byteio.ReadUint32BE(bin)
		switch err {
		case nil:
		case io.EOF:
			return nil
		default:
			return err
		}

		if err = byteio.WriteUint32LE(bout, n); err != nil {
			return err
		}
	}
}

func Example() {
	// prepare input buffer
	input := []uint32{0xDEADBEEF, 0x7FFFFFFF}
	inbuf := bytes.NewBuffer(nil)
	for _, n := range input {
		byteio.WriteUint32BE(inbuf, n)
	}

	// run our endianness swapper
	outbuf := bytes.NewBuffer(nil)
	if err := SwapEndian32(inbuf, outbuf); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// dump the output
	fmt.Println(hex.Dump(outbuf.Bytes()))

	// Output:
	// 00000000  ef be ad de ff ff ff 7f                           |........|
}
