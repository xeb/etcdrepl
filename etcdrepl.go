package main

import (
	"fmt"
	"github.com/xeb/etcdrepl/repl"
	"github.com/xeb/etcdrepl/third_party/github.com/codegangsta/cli"
	"github.com/xeb/etcdrepl/third_party/github.com/coreos/etcdctl/command"
)

func main() {
	app := repl.NewApp()
	app.Commands = []cli.Command{
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
		command.NewWatchCommand(),
		command.NewExecWatchCommand(),
		repl.NewMakeQuitCommand(),
	}
	e := app.Run()

	if e != nil {
		fmt.Println(e)
	}
}
