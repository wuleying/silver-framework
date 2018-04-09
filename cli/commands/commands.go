package commands

import (
	ucli "github.com/urfave/cli"
	"github.com/wuleying/silver-framework/globals"
)

var Commands = []ucli.Command{
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
