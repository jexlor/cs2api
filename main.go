package main

import (
	"log"
	"os"
	"time"

	"github.com/jexlor/cs2api/dev"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jexlor/cs2api/api"
	"github.com/jexlor/cs2api/db"
	"github.com/joho/godotenv"
)

//todo gracefully shutdown, add proper logging, env file managment of dev tools

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db, err := db.InitDB()
	if err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}
	handler := api.NewHandler(db)
	devhandler := dev.Devhandler(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to 8080 if not set
		log.Printf("No PORT specified, defaulting to %s", port)
	}

	router := setupRouter(handler, devhandler)

	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}

func setupRouter(handler *api.Handler, devhandler *dev.Handler) *gin.Engine {
	router := gin.Default()

	// Update CORS to allow the hx-trigger header and any other headers you may need
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // You can change this to the specific frontend URL if needed
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization", "hx-trigger"}, // Allow hx-trigger header
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.LoadHTMLGlob("templates/*")

	// Define your API routes
	apiGroup := router.Group("/cs2api")
	{
		apiGroup.GET("/", api.LandingPage)
		apiGroup.GET("/skins", handler.GetAllSkins)
		apiGroup.GET("/skins/search", handler.GetSkinById)
		apiGroup.GET("/skins/search/n", handler.GetSkinByName)
		apiGroup.GET("/collections", handler.GetCollections)
		apiGroup.GET("/collections/search/n", handler.GetCollectionByName)
		apiGroup.POST("/skins", devhandler.AddSkins)
		apiGroup.DELETE("/skins/delete", devhandler.DeleteSkinByName)
		apiGroup.PUT("/skins/edit", devhandler.UpdateSkinByName)
		apiGroup.PATCH("/skins/edit", devhandler.UpdateSkinByName)
	}
	//todo gracefully shutdown

	return router
}
