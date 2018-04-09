package main

import (
	ucli "github.com/urfave/cli"
	"github.com/wuleying/silver-framework/cli/commands"
	"github.com/wuleying/silver-framework/globals"
	"os"
)

func main() {
	run()
}

// Run 运行命令行
func run() {
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
	app.Commands = commands.Commands
	app.Run(os.Args)
}
