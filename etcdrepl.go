package main

import (
	"fmt"
	"github.com/xeb/etcdrepl/repl"
	"github.com/xeb/etcdrepl/third_party/github.com/codegangsta/cli"
	"github.com/xeb/etcdrepl/third_party/github.com/coreos/etcdctl/command"
	"os"
)

func main() {
	cmds := []cli.Command{
		command.NewMakeCommand(),
		command.NewMakeDirCommand(),
		command.NewRemoveCommand(),
		command.NewRemoveDirCommand(),
		command.NewGetCommand(),
		command.NewLsCommand(),
		command.NewSetCommand(),
		command.NewSetDirCommand(),
		command.NewUpdateCommand(),
		command.NewUpdateDirCommand(),

		// NOT YET SUPPORTED
		// command.NewWatchCommand(),
		// command.NewExecWatchCommand(),

		repl.NewMakeQuitCommand(),
	}
	app := repl.NewApp(cmds)
	e := app.Run(os.Stdin)

	if e != nil {
		fmt.Println(e)
	}
}
