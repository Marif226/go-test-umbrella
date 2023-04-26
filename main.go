package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	// Echo instance
	e := echo.New()

	e.Use(MyMiddleware)

	e.GET("/status", Handler)

	log.Fatal(e.Start(":8080"))
}

func Handler(ctx echo.Context) error {
	numberOfDays := time.Until(time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)).Hours() / 24
	err := ctx.String(http.StatusOK, fmt.Sprintf("%v days left\n", int(numberOfDays)))
	if err != nil {
		return err
	}

	return nil
}

func MyMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		val := ctx.Request().Header.Get("User-Role")
		if val == "admin" {
			log.Println("red button user detected")
		}

		err := next(ctx)
		if err != nil {
			return err
		}

		return nil
	}
}