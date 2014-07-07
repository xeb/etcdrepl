package repl

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/xeb/etcdrepl/third_party/github.com/codegangsta/cli"
	"io"
	"strings"
)

type App struct {
	Commands []cli.Command
	child    *cli.App
}

func NewApp() *App {

	ca := cli.NewApp()
	ca.Name = "etcdrepl"
	ca.Version = "0.0.1"
	ca.Flags = []cli.Flag{
		cli.BoolFlag{"debug", "output cURL commands which can be used to reproduce the request"},
		cli.BoolFlag{"no-sync", "don't synchronize cluster information before sending request"},
		cli.StringFlag{"output, o", "simple", "output response in the given format (`simple` or `json`)"},
		cli.StringSliceFlag{"peers, C", &cli.StringSlice{}, "a comma-delimited list of machine addresses in the cluster (default: {\"127.0.0.1:4001\"})"},
	}

	return &App{
		child: ca,
	}
}

func (a *App) Run(r io.Reader) error {

	if a == nil || a.Commands == nil || len(a.Commands) == 0 {
		return errors.New("No commands found")
	}

	for true {
		fmt.Print("etcdrepl> ")
		reader := Reader{bufio.NewReader(r)}
		text, e := reader.ReadStringAnyDelim('\n')

		if e != nil {
			return e
		}

		args := append([]string{"./etcdrepl"}, strings.Split(strings.Replace(text, "\n", "", 1000), " ")...)

		a.child.Commands = a.Commands

		e = a.child.Run(args)
		if e != nil {
			fmt.Printf("ERROR: %s\n", e)
		}
	}

	return nil
}
