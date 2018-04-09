package commands

import (
	ucli "github.com/urfave/cli"
)

var Commands = []ucli.Command{
	{
		Name:    "blockchain",
		Aliases: []string{"bc"},
		Usage:   "Blockchain opertaions",
		Subcommands: []ucli.Command{
			{
				Name:    "get",
				Aliases: []string{"g"},
				Usage:   "Get all blockchain info",
				Action: func(c *ucli.Context) error {
					return nil
				},
			},
		},
	},
}
