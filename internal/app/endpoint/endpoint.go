package endpoint

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
)

type Service interface {
	DaysLeft() int
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s,
	}
}

func (e *Endpoint) Status(ctx echo.Context) error {
	err := ctx.String(http.StatusOK, fmt.Sprintf("%v days left\n", e.s.DaysLeft()))
	if err != nil {
		return err
	}

	return nil
}