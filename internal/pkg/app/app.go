package app

import (
	"log"

	"github.com/Marif226/go-test-umbrella/internal/app/endpoint"
	"github.com/Marif226/go-test-umbrella/internal/app/middleware"
	"github.com/Marif226/go-test-umbrella/internal/app/service"
	"github.com/labstack/echo/v4"
)

type App struct {
	e 		*endpoint.Endpoint
	s 		*service.Service
	echo	*echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.s = service.New()

	a.e = endpoint.New(a.s)

	a.echo = echo.New()

	a.echo.Use(middleware.RoleChecker)

	a.echo.GET("/status", a.e.Status)

	return a, nil
}

func (a *App) Run() error {
	log.Println("server running...")

	err := a.echo.Start(":8080")
	if err != nil {
		return err
	}

	return nil
}