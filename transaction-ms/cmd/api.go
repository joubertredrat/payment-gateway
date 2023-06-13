package cmd

import (
	"fmt"
	"joubertredrat/transaction-ms/internal/infra"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
)

func getApiCommand() *cli.Command {
	return &cli.Command{
		Name:    "api",
		Aliases: []string{},
		Usage:   "Open HTTP api to listen",
		Action: func(c *cli.Context) error {
			config, err := infra.NewConfig()
			if err != nil {
				return err
			}

			r := gin.Default()
			if err := r.SetTrustedProxies(nil); err != nil {
				return err
			}

			statusController := infra.NewStatusController()

			r.NoRoute(func(c *gin.Context) {
				c.JSON(404, gin.H{"error": "404 page not found"})
			})

			ra := r.Group("/api")
			infra.RegisterCustomValidator()
			{
				ra.GET("/status", statusController.HandleStatus)
			}

			return r.Run(fmt.Sprintf("%s:%s", config.ApiHost, config.ApiPort))
		},
	}
}
