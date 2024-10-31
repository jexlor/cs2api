package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jexlor/cs2api/api"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Couldn't load env file!")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port must be set!")
	}

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/cs2api", api.LandingPage)
	fmt.Println("Running api on port:", port)

	err = router.Run(":" + port)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
