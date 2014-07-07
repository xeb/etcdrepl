package repl

import (
	"github.com/xeb/etcdrepl/third_party/github.com/codegangsta/cli"
	"testing"
)

type FakeStdin struct {
	Commands []string
}

func (f *FakeStdin) Read(p []byte) (n int, err error) {
	return 1, nil
}

func TestQuit(t *testing.T) {
	cmds := []cli.Command{
		NewMakeQuitCommand(),
	}
	app := NewApp(cmds)

	fr := &FakeStdin{
		Commands: []string{"help", "quit"},
	}

	_ = app.Run(fr)
}
