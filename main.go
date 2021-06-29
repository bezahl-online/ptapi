package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bezahl-online/ptapi/api"
	"github.com/bezahl-online/ptapi/mockserver"
	"github.com/bezahl-online/ptapi/param"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	swagger, err := api.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	serverPort, err := strconv.Atoi(strings.Split(swagger.Servers[0].URL, ":")[2])
	port := &serverPort
	e := echo.New()
	e.Use(middleware.CORS())

	var server api.ServerInterface
	if len(*param.TestDir) > 0 {
		fmt.Printf("starting in MOCK mode with test directory '%s'\n", *param.TestDir)
		server = &mockserver.API{}
	} else {
		server = &api.API{}
	}
	api.RegisterHandlers(e, server)

	if len(os.Getenv("PRODUCTIVE")) > 0 {
		e.Logger.Fatal(e.StartTLS(fmt.Sprintf("0.0.0.0:%d", *port), "localhost.crt", "localhost.key"))
	} else {
		e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%d", *port)))
	}

}
