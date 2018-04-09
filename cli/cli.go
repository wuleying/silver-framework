package cli

import (
	ucli "github.com/urfave/cli"
	"github.com/wuleying/silver-framework/cli/commands"
	"os"
)

// Run 运行命令行
func Run() {
	if len(os.Getenv("NODE_ID")) < 1 {
		os.Setenv("NODE_ID", "13000")
	}

	app := ucli.NewApp()
	app.Name = "Silver Framework"
	app.Usage = "CLI tools"
	app.Version = "0.1"
	app.Authors = []ucli.Author{
		ucli.Author{
			Name:  "Luo",
			Email: "lolooo@live.com",
		},
	}
	app.Commands = commands.Commands
	app.Run(os.Args)
}
