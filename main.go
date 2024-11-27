package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/jexlor/cs2api/db"
	"github.com/jexlor/cs2api/dev"

	"github.com/jexlor/cs2api/api"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db.InitDB()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to 8080 if not set
		log.Printf("No PORT specified, defaulting to %s", port)
	}

	router := setupRouter()

	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	// Change CORS as you wish
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.LoadHTMLGlob("templates/*")

	apiGroup := router.Group("/cs2api")
	{
		apiGroup.GET("/", api.LandingPage)
		apiGroup.GET("/skins", api.GetAllSkins)
		apiGroup.GET("/skins/search", api.GetSkinById)
		apiGroup.GET("/skins/search/n", api.GetSkinByName)
		apiGroup.GET("/collections", api.GetCollections)
		apiGroup.GET("/collections/search/n", api.GetCollectionByName)
		apiGroup.POST("/skins", dev.AddSkin)                   //hide for production
		apiGroup.DELETE("/skins/delete", dev.DeleteSkinByName) //hide for production
		apiGroup.PUT("/skins/edit", dev.UpdateSkinByName)      //hide for production
	}
	return router
}
