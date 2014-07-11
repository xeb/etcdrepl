package repl

import (
	"bytes"
	// "fmt"
	"github.com/xeb/etcdrepl/third_party/github.com/codegangsta/cli"
	"github.com/xeb/etcdrepl/third_party/github.com/coreos/etcdctl/command"
	"testing"
)

func TestHelp(t *testing.T) {
	cmds := []cli.Command{
		NewMakeQuitCommand(),
	}
	app := NewApp(cmds)

	_ = app.Run(bytes.NewReader([]byte("help\n")))

	// if e != nil {
	// 	fmt.Printf("Error %s", e)
	// 	t.FailNow()
	// }
}

func TestGet(t *testing.T) {
	cmds := []cli.Command{
		command.NewGetCommand(),
		NewMakeQuitCommand(),
	}
	app := NewApp(cmds)

	_ = app.Run(bytes.NewReader([]byte("get test\n")))

	// if e != nil {
	// 	fmt.Printf("Error %s", e)
	// 	t.FailNow()
	// }
}
