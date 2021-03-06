package command

import (
	"fmt"

	"github.com/xeb/etcdrepl/third_party/github.com/codegangsta/cli"
	"github.com/xeb/etcdrepl/third_party/github.com/coreos/go-etcd/etcd"
)

func NewLsCommand() cli.Command {
	return cli.Command{
		Name:	"ls",
		Usage:	"retrieve a directory",
		Flags: []cli.Flag{
			cli.BoolFlag{"recursive", "returns all values for key and child keys"},
		},
		Action: func(c *cli.Context) {
			handleLs(c, lsCommandFunc)
		},
	}
}

// handleLs handles a request that intends to do ls-like operations.
func handleLs(c *cli.Context, fn handlerFunc) {
	handlePrint(c, fn, printLs)
}

// printLs writes a response out in a manner similar to the `ls` command in unix.
// Non-empty directories list their contents and files list their name.
func printLs(resp *etcd.Response, format string) {
	if !resp.Node.Dir {
		fmt.Println(resp.Node.Key)
	}
	for _, node := range resp.Node.Nodes {
		rPrint(&node)
	}
}

// lsCommandFunc executes the "ls" command.
func lsCommandFunc(c *cli.Context, client *etcd.Client) (*etcd.Response, error) {
	key := "/"
	if len(c.Args()) != 0 {
		key = c.Args()[0]
	}
	recursive := c.Bool("recursive")

	// Retrieve the value from the server.
	return client.Get(key, false, recursive)
}

// rPrint recursively prints out the nodes in the node structure.
func rPrint(n *etcd.Node) {
	fmt.Println(n.Key)
	for _, node := range n.Nodes {
		rPrint(&node)
	}
}
