package handler

import (
	lv "dfc/public/view/league"
	pv "dfc/public/view/player"
	"dfc/service"
	"dfc/types"

	"github.com/labstack/echo/v4"
)

type PlayerService interface {
	ListPlayersForLeague() ([]service.Player, error)
	CreatePlayer(types.PlayerParams) (service.Player, error)
	GetPlayer(string) (service.Player, error)
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
	players, err := ph.PlayerServices.ListPlayersForLeague()
	if err != nil {
		return err
	}

	return Render(c, lv.Index(players))
}

func (ph *PlayerHandler) NewPlayerHandler(c echo.Context) error {
	return Render(c, pv.NewPlayer())
}

func (ph *PlayerHandler) ShowPlayerHandler(c echo.Context) error {
	playerID := c.QueryParam("id")
	player, _ := ph.PlayerServices.GetPlayer(playerID)
	return Render(c, pv.ShowPlayer(player))
}

func (ph *PlayerHandler) CreatePlayerHandler(c echo.Context) error {
	params := types.PlayerParams{
		Name: c.FormValue("name"),
		Deck: c.FormValue("deck"),
	}

	_, err := ph.PlayerServices.CreatePlayer(params)
	if err != nil {
		return Render(c, pv.NewPlayerWithErrors(params, types.PlayerErrors{}))
	}

	return c.Redirect(302, "/")
}
