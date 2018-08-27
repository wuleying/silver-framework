package commands

import ucli "github.com/urfave/cli"

var CommandList = []ucli.Command{
	{
		Name:    "version",
		Aliases: []string{"V"},
		Usage:   "Print version info",
		Action: func(c *ucli.Context) error {
			return PrintVersion()
		},
	},
}
