package cmd

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go.elastic.co/apm/module/apmechov4/v2"
	"net/http"
	"svc-activity/config"
	api "svc-activity/internal/adapter/handler/http/handlers"
	"svc-activity/internal/adapter/handler/http/routes"
	"svc-activity/utils"
)

var greetHTTP = `
     _        _   _       _ _           _   _ _____ _____ ____  
    / \   ___| |_(_)_   _(_) |_ _   _  | | | |_   _|_   _|  _ \ 
   / _ \ / __| __| \ \ / / | __| | | | | |_| | | |   | | | |_) |
  / ___ \ (__| |_| |\ V /| | |_| |_| | |  _  | | |   | | |  __/ 
 /_/   \_\___|\__|_| \_/ |_|\__|\__, | |_| |_| |_|   |_| |_|    
                                |___/                           
`

func RegisterHTTP() *cobra.Command {
	return &cobra.Command{
		Use:   "http",
		Short: "activity http",
		Run: func(cmd *cobra.Command, args []string) {

			// greet
			fmt.Println(greetHTTP)

			// http
			startHttp()

		},
	}
}

func startHttp() {

	// init
	e := echo.New()

	// middlewares
	e.Use(middleware.Recover())
	e.Use(apmechov4.Middleware())

	// check health
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello from the other side!")
	})

	// setup config
	getConfig, err := config.SetupConfig()
	if err != nil {
		logrus.Fatal(fmt.Sprintf("failed to set up configuration err: %v", err.Error()))
	}

	// init injector
	inject := config.InitInjection(getConfig)
	handler := api.NewHandler(inject)

	// register route
	routes.RegisterRoute(e, handler)

	// start http
	go func() {
		if err := e.Start(fmt.Sprintf(":%s", getConfig.App.Port)); err != http.ErrServerClosed {
			logrus.Fatal(fmt.Sprintf("failed to set up http server err: %v", err.Error()))
		}
	}()

	// terminate signal
	utils.WaitTerminateSignal()

	// close config
	getConfig.CloseConfig()
}
