package handlers

import (
	"dfc/public/view"
	"dfc/services"

	"github.com/labstack/echo/v4"
)

type PlayerService interface {
	ListPlayersForSeason() ([]services.Player, error)
}

func NewPlayerHandler(ps PlayerService) *PlayerHandler {
	return &PlayerHandler{
		PlayerServices: ps,
	}
}

type PlayerHandler struct {
	PlayerServices PlayerService
}

func (ph *PlayerHandler) LeaderboardHandler(c echo.Context) error {
	players, err := ph.PlayerServices.ListPlayersForSeason()
	if err != nil {
		return err
	}

	component := view.Index(players)
	return component.Render(c.Request().Context(), c.Response().Writer)
}
