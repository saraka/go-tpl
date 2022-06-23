package cmd

import (
	"github.com/saraka/go-tpl/internal/app/application"
	"github.com/urfave/cli/v2"
)

var Add = cli.Command{
	Name:    "add",
	Aliases: []string{"i"},
	Usage:   "add a new application",
	Action:  addAction,
}

func addAction(c *cli.Context) (err error) {
	return application.Exec(c.Args().First())
}

func init() {
	Commands = append(Commands, &Add)
}
