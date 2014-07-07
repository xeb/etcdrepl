package repl

import (
	"bufio"
)

type Reader struct {
	*bufio.Reader
}

func (b *Reader) ReadStringAnyDelim(delim byte) (line string, err error) {
	bytes, err := b.ReadBytes(delim)
	line = string(bytes)
	return line, err
}
