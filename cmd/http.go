package cmd

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go-hexa/config"
	_ "go-hexa/docs"
	api "go-hexa/internal/adapter/handler/http/handlers"
	"go-hexa/internal/adapter/handler/http/routes"
	"go-hexa/utils"
	"go.elastic.co/apm/module/apmfiber/v2"
)

var greetHTTP = `
 _   _ _____ _____ ____     ____ __  __ ____  
| | | |_   _|_   _|  _ \   / ___|  \/  |  _ \ 
| |_| | | |   | | | |_) | | |   | |\/| | | | |
|  _  | | |   | | |  __/  | |___| |  | | |_| |
|_| |_| |_|   |_| |_|      \____|_|  |_|____/ 
											  
`

func RegisterHTTP() *cobra.Command {
	return &cobra.Command{
		Use:   "http",
		Short: "http entrypoint",
		Run: func(cmd *cobra.Command, args []string) {

			// greet
			fmt.Println(greetHTTP)

			// http
			startHttp()

		},
	}
}

// @Title			Golang Hexa Swagger
// @Version			1.0
// @Description		Golang hexagonal swagger documentation
// @Contact.name	Adam Syarif Hidayatullah
// @Contact.email	adamsyarif219@gmail.com
// @Host			/go-hexa
// @Schemes			http https
// @BasePath		/
func startHttp() {

	// init
	e := fiber.New()

	// middlewares
	registerGlobalMiddlewares(e)

	// check health
	e.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello from the other side!")
	})

	// swagger
	e.Get("/docs/*", swagger.HandlerDefault)

	// setup config
	getConfig, err := config.SetupConfig()
	if err != nil {
		logrus.Fatal(utils.PrintMessageWithError("failed to set up configuration", err))
	}

	// init injector
	inject := config.InitInjection(getConfig)
	handler := api.NewHandler(inject)

	// register route
	routes.RegisterRoute(e, handler, getConfig.ApiKey)

	// start http
	go func() {
		if err := e.Listen(fmt.Sprintf(":%s", getConfig.App.Port)); err != nil {
			logrus.Fatal(utils.PrintMessageWithError("failed to set up http server", err))
		}
	}()

	// terminate signal
	utils.WaitTerminateSignal()

	// close config
	getConfig.CloseConfig()
}

func registerGlobalMiddlewares(e *fiber.App) {
	e.Use(recover.New())
	e.Use(apmfiber.Middleware())
	e.Use(logger.New())
	e.Use(pprof.New())
	e.Use(healthcheck.New())
	e.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders: "*",
	}))
}
