package main

import (
	"dfc/db"
	"dfc/handlers"
	"dfc/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	DB_NAME string = "data/dfc.db"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	store, err := db.NewStore(DB_NAME)
	if err != nil {
		e.Logger.Fatalf("failed to create store: %s", err)
	}

	playerService := services.NewPlayerServices(services.Player{}, store)
	playerHandler := handlers.NewPlayerHandler(playerService)

	e.Static("/assets", "assets")
	e.GET("/", playerHandler.LeaderboardHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
