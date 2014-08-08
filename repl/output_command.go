package repl

import (
	// "fmt"
	"github.com/xeb/etcdrepl/third_party/github.com/codegangsta/cli"
)

func NewMakeOutputCommand() cli.Command {
	return cli.Command{
		Name:  "output",
		Usage: "output response in the given format (`simple` or `json`)",
		Action: func(c *cli.Context) {
			// No action
			// output := c.Args()[0]
			// for _, flag := range c.App.Flags {
			// 	if flag.getName() == "output" {
			// 		fmt.Printf("Found it")
			// 	}
			// }
		},
	}
}
