package main

import (
	ucli "github.com/urfave/cli"
	"github.com/wuleying/silver-framework/cli/commands"
	"github.com/wuleying/silver-framework/version"
	"os"
)

func main() {
	run()
}

// run 运行命令行
func run() {
	app := ucli.NewApp()
	app.Name = "Silver Framework"
	app.Usage = "CLI tools"
	app.Version = version.Version
	app.Authors = []ucli.Author{
		ucli.Author{
			Name:  "Silver",
			Email: "lolooo@live.com",
		},
	}
	app.Commands = commands.CommandList
	app.Run(os.Args)
}
