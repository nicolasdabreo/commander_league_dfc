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

type Result struct {
	gorm.Model
	PodSize                int  `gorm:"not null;"`
	Place                  int  `gorm:"not null;"`
	TheCouncilOfWizards    bool `gorm:"default:false"`
	DavidAndTheGoliaths    bool `gorm:"default:false"`
	Untouchable            bool `gorm:"default:false"`
	Cleave                 bool `gorm:"default:false"`
	ItsFreeRealEstate      bool `gorm:"default:false"`
	IAmTimmy               bool `gorm:"default:false"`
	BigBiggerHuge          bool `gorm:"default:false"`
	CloseButNoCigar        bool `gorm:"default:false"`
	JustAsGarfieldIntended bool `gorm:"default:false"`
}

func init() {
	var err error

	db, err = gorm.Open("sqlite3", "data/dfc.db")
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Player{})
	db.AutoMigrate(&Result{})
}

func main() {
	defer db.Close()

	router := gin.Default()

	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", leaderboardHandler)
	router.GET("/players/:id", showPlayerHandler)
	router.GET("/players/new", newPlayerHandler)
	router.POST("/players", createPlayerHandler)
	router.GET("/results/new", newResultHandler)
	router.POST("/results", createResultHandler)
	router.GET("/downloads/rulepack", rulepackHandler)

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

func showPlayerHandler(c *gin.Context) {
	playerID := c.Param("id")
	var player Player
	result := db.Where("id = ?", playerID).First(&player)

	if result.Error != nil {
		c.HTML(404, "not_found.tmpl.html", nil)
		return
	}

	c.HTML(200, "show_player.tmpl.html", gin.H{"Player": result.Value})
}

func createResultHandler(c *gin.Context) {
	// Parse form data
	var form struct {
		PodSize                int  `form:"pod_size" binding:"required"`
		Place                  int  `form:"place" binding:"required"`
		TheCouncilOfWizards    bool `form:"the_council_of_wizards"`
		DavidAndTheGoliaths    bool `form:"david_and_the_goliaths"`
		Untouchable            bool `form:"untouchable"`
		Cleave                 bool `form:"cleave"`
		ItsFreeRealEstate      bool `form:"its_free_real_estate"`
		IAmTimmy               bool `form:"i_am_timmy"`
		BigBiggerHuge          bool `form:"big_bigger_huge"`
		CloseButNoCigar        bool `form:"close_but_no_cigar"`
		JustAsGarfieldIntended bool `form:"just_as_garfield_intended"`
	}

	if err := c.ShouldBind(&form); err != nil {
		c.HTML(http.StatusBadRequest, "new_result.tmpl.html", gin.H{"error": err.Error()})
		return
	}

	// Create a new result
	newResult := Result{
		Place:                  form.Place,
		TheCouncilOfWizards:    form.TheCouncilOfWizards,
		DavidAndTheGoliaths:    form.DavidAndTheGoliaths,
		Untouchable:            form.Untouchable,
		Cleave:                 form.Cleave,
		ItsFreeRealEstate:      form.ItsFreeRealEstate,
		IAmTimmy:               form.IAmTimmy,
		BigBiggerHuge:          form.BigBiggerHuge,
		CloseButNoCigar:        form.CloseButNoCigar,
		JustAsGarfieldIntended: form.JustAsGarfieldIntended,
	}

	if err := db.Create(&newResult).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "new_result.tmpl.html", gin.H{"error": "Failed to create result"})
		return
	}

	// Redirect to the index page or another success page
	c.Redirect(http.StatusSeeOther, "/")
}

func rulepackHandler(c *gin.Context) {
	targetPath := "./assets/images/commander_league_rulepack.pdf"

	//Seems this headers needed for some browsers (for example without this headers Chrome will download files as txt)
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename=commander_league_rulepack.pdf")
	c.Header("Content-Type", "application/octet-stream")

	c.FileAttachment(targetPath, "commander_league_rulepack.pdf")
}
