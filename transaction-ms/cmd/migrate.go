package cmd

import (
	"joubertredrat/transaction-ms/internal/infra"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/urfave/cli/v2"
)

func getMigrateCommand() *cli.Command {
	return &cli.Command{
		Name:    "migrate",
		Aliases: []string{},
		Usage:   "Run migration into database",
		Action: func(c *cli.Context) error {
			config, err := infra.NewConfig()
			if err != nil {
				return err
			}

			db, err := infra.GetDatabaseConnection(infra.GetMysqlDSN(
				config.DatabaseHost,
				config.DatabasePort,
				config.DatabaseName,
				config.DatabaseUser,
				config.DatabasePassword,
			))
			if err != nil {
				return err
			}

			driver, err := mysql.WithInstance(db, &mysql.Config{})
			if err != nil {
				return err
			}

			m, err := migrate.NewWithDatabaseInstance("file://migrations", "mysql", driver)
			if err != nil {
				return err
			}

			return m.Up()
		},
	}
}
