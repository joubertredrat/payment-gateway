package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func getMigrateCommand() *cli.Command {
	return &cli.Command{
		Name:    "migrate",
		Aliases: []string{},
		Usage:   "Run migration into database",
		Action: func(c *cli.Context) error {
			fmt.Println("run migration")
			return nil
		},
	}
}
