package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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
