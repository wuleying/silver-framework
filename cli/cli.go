package main

import (
	ucli "github.com/urfave/cli"
	"github.com/wuleying/silver-framework/globals"
	"os"
)

func main() {
	run()
}

// Run 运行命令行
func run() {
	var commandList = []ucli.Command{
		{
			Name:    "version",
			Aliases: []string{"V"},
			Usage:   "Print version info",
			Action: func(c *ucli.Context) error {
				println(globals.Version)
				return nil
			},
		},
	}

	app := ucli.NewApp()
	app.Name = "Silver Framework"
	app.Usage = "CLI tools"
	app.Version = globals.Version
	app.Authors = []ucli.Author{
		ucli.Author{
			Name:  "Luo",
			Email: "lolooo@live.com",
		},
	}
	app.Commands = commandList
	app.Run(os.Args)
}
