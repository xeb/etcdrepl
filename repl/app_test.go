package repl

import (
	"github.com/xeb/etcdrepl/third_party/github.com/codegangsta/cli"
	"os"
	"repl"
	"testing"
)

type FakeStdin struct {
	Commands []string
}

func (f *FakeStdin) Read(p []byte) (n int, err error) {
	return 1, nil
}

func TestQuit(t *testing.T) {
	app := repl.NewApp()
	app.Commands = []cli.Command{
		repl.NewMakeQuitCommand(),
	}

	fr := &FakeStdin{
		Commands: ["test","quit"],
	}

	e := app.Run(fr)
}
