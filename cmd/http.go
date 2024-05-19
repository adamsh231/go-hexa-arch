package cmd

import (
	"fmt"
	"net/http"
	"svc-activity/config"
	api "svc-activity/internal/adapter/handler/http/handlers"
	"svc-activity/internal/adapter/handler/http/routes"
	"svc-activity/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.elastic.co/apm/module/apmechov4/v2"

	_ "svc-activity/docs"
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

// @Title			Majoo Logging Activity
// @Version		1.0
// @Description	Majoo Logging Activity
// @Contact.name	Adam Syarif Hidayatullah
// @Contact.email	adam@majoo.id
// @Host			/svc-activity
// @Schemes		http https
// @BasePath		/
func startHttp() {

	// init
	e := echo.New()

	// middlewares
	registerGlobalMiddlewares(e)

	// check health
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello from the other side!")
	})

	// swagger
	e.GET("/docs/*", echoSwagger.WrapHandler)

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
		if err := e.Start(fmt.Sprintf(":%s", getConfig.App.Port)); err != http.ErrServerClosed {
			logrus.Fatal(utils.PrintMessageWithError("failed to set up http server", err))
		}
	}()

	// terminate signal
	utils.WaitTerminateSignal()

	// close config
	getConfig.CloseConfig()
}

func registerGlobalMiddlewares(e *echo.Echo) {
	e.Use(middleware.Recover())
	e.Use(apmechov4.Middleware())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
}
