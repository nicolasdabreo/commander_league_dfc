package main

import (
	"embed"
	"fmt"
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
	db, err = gorm.Open("sqlite3", "test.db")
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

	router.GET("/", renderIndex)

	// Define your routes here

	router.Run(":8080")
}

func renderIndex(c *gin.Context) {
	c.HTML(200, "index.tmpl.html", gin.H{})
}
