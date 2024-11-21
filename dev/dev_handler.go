package dev

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jexlor/cs2api/api"
	"github.com/jexlor/cs2api/db"
)

// this package is for development only.
// uncomment handler functions here and endpoints in main.go file (it's labeled as "endpoints")
// (please maintain view only behavior for api if you are not a developer)

func AddSkin(c *gin.Context) {
	var skin api.Skin

	if err := c.ShouldBindJSON(&skin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't parse json!"})
		return
	}

	_, err := db.DB.Exec(`INSERT INTO skins(name, rarity, collection, quality, price, url)
	VALUES ($1, $2, $3, $4,$5, $6)`, skin.Name, skin.Rarity, skin.Collection, skin.Quality, skin.Price, skin.Url)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't add skin!"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Skin added!"})
}
