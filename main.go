package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jexlor/cs2api/db"
	"github.com/jexlor/cs2api/dev"

	"github.com/jexlor/cs2api/api"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	db.InitDB()
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Couldn't load env file!")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port must be set!")
	}

	router := gin.Default()

	//endpoints
	router.LoadHTMLGlob("templates/*")
	router.GET("/cs2api", api.LandingPage)
	router.GET("/cs2api/skins", api.GetAllSkins)
	router.GET("/cs2api/skins/search", api.GetSkinById)
	router.GET("/cs2api/skins/search/n", api.GetSkinByName)
	router.GET("/cs2api/collections", api.GetCollections)
	router.GET("/cs2api/collections/search/n", api.GetCollectionByName)
	router.POST("/cs2api/skins", dev.AddSkin)
	fmt.Println("Running api on port:", port)

	err = router.Run(":" + port)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
