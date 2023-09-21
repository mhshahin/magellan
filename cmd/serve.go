package cmd

import (
	"github.com/labstack/echo/v4"
	"github.com/mhshahin/magellan/config"
	"github.com/mhshahin/magellan/handlers"
	"github.com/mhshahin/magellan/repository"
	"github.com/mhshahin/magellan/routes"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "This command serves the APIs",
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func serve() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	redisClient, err := config.InitializeRedis(cfg.Redis)
	if err != nil {
		panic(err)
	}

	repos := repository.NewRepository(redisClient)
	handlers := handlers.NewHandlers(repos)

	e := echo.New()

	routes.InitializeRoutes(e, handlers)

	e.Logger.Fatal(e.Start(":7080"))
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
