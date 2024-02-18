package handler

import (
	"context"
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

	if errs, hasErrs := params.Validate(); hasErrs {
		setFlash(c, "error", "Check form for errors")
		return Render(c, pv.NewPlayerWithErrors(params, errs))
	}

	if _, err := ph.PlayerServices.CreatePlayer(params); err != nil {
		return Render(c, pv.NewPlayer())
	}

	return c.Redirect(302, "/")
}

func setFlash(c echo.Context, typ string, err string) {
	ctx := c.Request().Context()
	ctx = ctxWithFlash(ctx, typ, err)
	c.SetRequest(c.Request().WithContext(ctx))
}

func ctxWithFlash(ctx context.Context, typ string, err string) context.Context {
	return context.WithValue(ctx, typ, err)
}
