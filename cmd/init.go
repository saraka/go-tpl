package cmd

import (
	"github.com/saraka/go-tpl/internal/app/initialization"
	"github.com/urfave/cli/v2"
)

var Init = cli.Command{
	Name:    "init",
	Aliases: []string{"i"},
	Usage:   "Initialize a new project",
	Action:  initAction,
}

func initAction(c *cli.Context) (err error) {
	return initialization.Init(c.Args().First())
}
