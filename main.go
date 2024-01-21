package main

import (
	"embed"
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//go:embed templates/*
var resources embed.FS

var db *gorm.DB

var t = template.Must(template.ParseFS(resources, "templates/*"))

type Player struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
	Deck string `gorm:"not null"`
}

func init() {
	var err error
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Player{}) // replace YourModel with your actual model
}

func main() {
	defer db.Close()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, Gin!"})
	})

	// Define your routes here

	router.Run(":8080")
}
