package dev

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jexlor/cs2api/api"
	"github.com/jexlor/cs2api/db"
)

// this package is for development only, hide this handlers and remove endpoints in main.go file for production.

func AddSkins(c *gin.Context) {
	var skins []api.Skin

	if err := c.ShouldBindJSON(&skins); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't parse JSON!"})
		return
	}

	stmt, err := db.DB.Prepare(`INSERT INTO skins(name, weapon, rarity, collection,  price, stattrack_price, url)
	VALUES ($1, $2, $3, $4, $5, $6, $7)`)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't prepare SQL statement!"})
		return

	}
	defer stmt.Close()

	for _, skin := range skins {
		_, err := stmt.Exec(skin.Name, skin.Weapon, skin.Rarity, skin.Collection, skin.Price, skin.StattrackPrice, skin.Url)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't add one or more skins!"})
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{"message": "All skins added successfully!"})
}

func DeleteSkinByName(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name parameter is required!"})
		return
	}

	err := DeleteSkinByNameJson(name)
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

	err := UpdateSkinByNameJson(name, updatedSkin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't update skin!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Skin updated successfully!"})
}
