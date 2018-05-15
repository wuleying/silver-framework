package main

import (
	"os"
	ucli "github.com/urfave/cli"
	"github.com/wuleying/silver-framework/globals"
	"github.com/wuleying/silver-framework/cli/commands"
)

func main() {
	run()
}

// run 运行命令行
func run() {
	var commandList = []ucli.Command{
		{
			Name:    "version",
			Aliases: []string{"V"},
			Usage:   "Print version info",
			Action: func(c *ucli.Context) error {
				return commands.PrintVersion()
			},
		},
	}

	app := ucli.NewApp()
	app.Name = "Silver Framework"
	app.Usage = "CLI tools"
	app.Version = globals.Version
	app.Authors = []ucli.Author{
		ucli.Author{
			Name:  "Silver",
			Email: "lolooo@live.com",
		},
	}
	app.Commands = commandList
	app.Run(os.Args)
}
