package main

import (
	"embed"
	"fmt"
	"log"
	"os"

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

	absPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("Absolute Path:", absPath)

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", renderPlayerLeaderboard)
	router.GET("/players/new", renderNewPlayerForm)
	router.POST("/players", createPlayer)
	router.GET("/results/new", renderNewResultForm)

	// Define your routes here

	router.Run(":8080")
}

func renderPlayerLeaderboard(c *gin.Context) {
	c.HTML(200, "player_leaderboard.tmpl.html", gin.H{})
}

func renderNewPlayerForm(c *gin.Context) {
	c.HTML(200, "new_player.tmpl.html", gin.H{})
}

func renderNewResultForm(c *gin.Context) {
	c.HTML(200, "new_result.tmpl.html", gin.H{})
}

func createPlayer(c *gin.Context) {
	var player Player

	log.Println(player)

	if err := c.ShouldBindJSON(&player); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Create the player in the database
	if err := db.Create(&player).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create player"})
		return
	}

	c.JSON(201, player)
}
