package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jexlor/cs2api/db"
)

func LandingPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
func GetAllSkins(c *gin.Context) {
	skins, err := getAllSkinsJson()
	if err != nil {
		log.Printf("Error fetching skins: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't serve skins!"})
		return
	}

	c.JSON(http.StatusOK, skins)
}

func AddSkin(c *gin.Context) {
	var skin Skin

	if err := c.ShouldBindJSON(&skin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't parse json!"})
		return
	}

	_, err := db.DB.Exec(`INSERT INTO skins(name, rarity, collection, quality, price, url)
	VALUES (?, ?, ?, ?, ?, ?)`, skin.Name, skin.Rarity, skin.Collection, skin.Quality, skin.Price, skin.Url)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't add skin!"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Skin added!"})
}
