package main

import (
	"embed"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//go:embed templates/*
var resources embed.FS

var db *gorm.DB

type Player struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
	Deck string `gorm:"not null"`
}

func init() {
	var err error

	db, err = gorm.Open("sqlite3", "dfc.db")
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Player{})
}

func main() {
	defer db.Close()

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", leaderboardHandler)
	router.GET("/players/new", newPlayerHandler)
	router.POST("/players", createPlayerHandler)
	router.GET("/results/new", newResultHandler)

	// Define your routes here

	router.Run(":8080")
}

func leaderboardHandler(c *gin.Context) {
	var players []Player
	results := db.Find(&players)

	c.HTML(200, "player_leaderboard.tmpl.html", gin.H{"Players": results.Value})
}

func newPlayerHandler(c *gin.Context) {
	c.HTML(200, "new_player.tmpl.html", gin.H{})
}

func newResultHandler(c *gin.Context) {
	c.HTML(200, "new_result.tmpl.html", gin.H{})
}

func createPlayerHandler(c *gin.Context) {
	// Parse form data
	var form struct {
		Name string `form:"name" binding:"required"`
		Deck string `form:"deck" binding:"required,oneof=tokens dragons goad flying zombies"`
	}

	if err := c.ShouldBind(&form); err != nil {
		c.HTML(http.StatusBadRequest, "new_player.tmpl.html", gin.H{"error": err.Error()})
		return
	}

	// Create a new player
	newPlayer := Player{Name: form.Name, Deck: form.Deck}
	if err := db.Create(&newPlayer).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "new_player.tmpl.html", gin.H{"error": "Failed to create player"})
		return
	}

	// Redirect to the index page or another success page
	c.Redirect(http.StatusSeeOther, "/")
}
