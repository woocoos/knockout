package main

import (
	"github.com/urfave/cli/v2"
	"github.com/woocoos/knockout/cmd/tools/res"
	"log"
	"os"
)

const Version = "0.0.1"

var commands = []*cli.Command{
	res.Cmd,
}

func main() {
	app := cli.NewApp()
	app.Name = "kocli"
	app.Usage = "a cli command for knockout project"
	app.Version = Version
	app.Commands = commands
	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err.Error())
	}
}
