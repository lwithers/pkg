/*
Package byteio supports applications that perform byte-oriented I/O, reading and
writing only small chunks at a time. Its primary purpose is the Reader and
Writer interfaces which extend io.Reader and io.Writer with byte-, rune- and
string-oriented functions, along with an adapter function that turns any reader
or writer into the interface (possibly by wrapping it in a bufio.Reader or
bufio.Writer).

Note that the Reader and Writer types from both the bufio and the bytes package
implement this package's Reader and Writer interfaces.

When using the adapter function for the writer, since it is possible that the
returned type may be a new bufio.Writer, it is necessary to check for whether
Flush() must be called at operation completion time. This can be done
succinctly with:

	bout := byteio.NewWriter(out)
	defer byteio.FlushIfNecessary(out)

The binary read/write functions were benchmarked (using bufio) to determine that,
for sizes up to and including 8 bytes, it was faster to call
ReadByte()/WriteByte() multiple times in succession than to call Read()/Write()
with a small buffer.
*/
package byteio
