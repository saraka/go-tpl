package main

import (
	"log"
	"os"

	"github.com/saraka/go-tpl/cmd"
	"github.com/urfave/cli/v2"
)

const (
	AppName string = "gotpl"
)

func main() {
	var app = &cli.App{
		Usage:    "\b\b ",
		Commands: cmd.Commands,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
