package main

import (
	"dfc/db"
	"dfc/handler"
	"dfc/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	DB_NAME string = "_data/dfc.db"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	store, err := db.NewStore(DB_NAME)
	if err != nil {
		e.Logger.Fatalf("failed to create store: %s", err)
	}

	playerService := service.NewPlayerServices(service.Player{}, store)
	playerHandler := handler.NewPlayerHandler(playerService)

	resultService := service.NewResultServices(service.Result{}, store)
	resultHandler := handler.NewResultHandler(resultService)

	e.Static("/assets", "assets")

	e.GET("/", playerHandler.LeaderboardHandler)
	e.GET("/players/new", playerHandler.NewPlayerHandler)
	e.GET("/players/:id", playerHandler.ShowPlayerHandler)
	e.POST("/players", playerHandler.CreatePlayerHandler)
	e.GET("/results/new", resultHandler.NewResultHandler)
	// e.POST("/results", resultHandler.NewResultHandler)
	// e.GET("/downloads/rulepack", fileHandler.rulepackHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
