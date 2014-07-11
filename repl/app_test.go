package repl

import (
	// "bufio"
	"errors"
	"fmt"
	"github.com/xeb/etcdrepl/third_party/github.com/codegangsta/cli"
	"testing"
)

var EOF = errors.New("EOF")

type FakeStdin struct {
	Commands  []string
	lastIndex int
}

func NewFakeStdin(cmds []string) (f *FakeStdin) {
	return &FakeStdin{
		Commands:  cmds,
		lastIndex: 0,
	}
}

func (f *FakeStdin) Read(p []byte) (n int, err error) {

	if len(f.Commands) >= f.lastIndex {
		return 0, EOF
	}

	cmd := f.Commands[f.lastIndex]
	fmt.Printf("Command is %s ", cmd)
	f.lastIndex = f.lastIndex + 1

	if len(p) > len(cmd) {
		p = []byte(cmd)
		return len(cmd), nil
	} else {
		p = []byte(cmd[0:len(p)])
		return len(p), nil
	}
}

func TestQuit(t *testing.T) {
	cmds := []cli.Command{
		NewMakeQuitCommand(),
	}
	app := NewApp(cmds)

	fr := &FakeStdin{
		Commands: []string{"help\n", "quit\n"},
	}

	_ = app.Run(fr)
}

// func TestReader(t *testing.T) {

// 	fr := NewFakeStdin([]string{"help.me\n", "quit.oh\n"})

// 	reader := Reader{bufio.NewReader(fr)}
// 	fmt.Printf("----WOOT\n\n r is %s\nreader is %s", fr, reader)
// 	text, e := reader.ReadStringAnyDelim([]byte{'.'})

// 	fmt.Printf("Text is %s (error is %s)", text, e)
// }
