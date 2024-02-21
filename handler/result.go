package handler

import (
	rv "dfc/public/view/result"

	"github.com/labstack/echo/v4"
)

type ResultService interface {
}

func NewResultHandler(ps ResultService) *ResultHandler {
	return &ResultHandler{
		ResultServices: ps,
	}
}

type ResultHandler struct {
	ResultServices ResultService
}

func (ph *ResultHandler) NewResultHandler(c echo.Context) error {
	return Render(c, rv.NewResult())
}
