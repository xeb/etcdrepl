package repl

import (
	"github.com/xeb/etcdrepl/third_party/github.com/codegangsta/cli"
	"os"
)

func NewMakeQuitCommand() cli.Command {
	return cli.Command{
		Name:  "quit",
		Usage: "exits the current program",
		Action: func(c *cli.Context) {
			os.Exit(0)
		},
	}
}
