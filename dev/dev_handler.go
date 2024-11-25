package dev

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jexlor/cs2api/api"
	"github.com/jexlor/cs2api/db"
)

// this package is for development only, hide this handlers and remove endpoints in main.go file for production.

func AddSkin(c *gin.Context) {
	var skin api.Skin

	if err := c.ShouldBindJSON(&skin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't parse json!"})
		return
	}

	_, err := db.DB.Exec(`INSERT INTO skins(name, rarity, collection, quality, price, url)
	VALUES ($1, $2, $3, $4, $5, $6)`, skin.Name, skin.Rarity, skin.Collection, skin.Quality, skin.Price, skin.Url)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't add skin!"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Skin added!"})
}

func DeleteSkinByName(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name parameter is required!"})
		return
	}

	err := deleteSkinByNameJson(name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't delete skin!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": "Skin deleted successfully!"})
}

func UpdateSkinByName(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name parameter is required!"})
		return
	}

	var updatedSkin api.Skin
	if err := c.ShouldBindJSON(&updatedSkin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body!"})
		return
	}

	err := updateSkinByNameJson(name, updatedSkin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't update skin!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Skin updated successfully!"})
}
