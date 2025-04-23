package api

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/jexlor/cs2api/db"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	db *db.Database
}

func NewHandler(db *db.Database) *Handler {
	return &Handler{db: db}
}

func LandingPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
func (h *Handler) GetAllSkins(c *gin.Context) {
	skins, err := GetAllSkinsJson(h.db)
	if err != nil {
		log.Printf("Error fetching skins: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't serve skins!"})
		return
	}

	c.JSON(http.StatusOK, skins)
}
func (h *Handler) GetSkinById(c *gin.Context) {
	idParam := c.Query("id")
	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id parameter is required!"})
		return
	}

	Id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Id parameter!"})
		return
	}

	skin, err := GetSkinByIdJson(h.db, Id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Couldn't find skin with that id!"})
		return
	}

	c.JSON(http.StatusOK, skin)
}

func (h *Handler) GetSkinByName(c *gin.Context) {
	name := c.Query("name")
	name = strings.TrimSpace(name)
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name parameter is required!"})
		return
	}

	skin, err := GetSkinByNameJson(h.db, name)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Couldn't find skin with that name!"})
		return
	}
	c.JSON(http.StatusOK, skin)
}
func (h *Handler) GetCollectionByName(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	name = strings.TrimSpace(name)
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name parameter is required!"})
		return
	}

	skinsFromCollection, err := GetCollectionByNameJson(h.db, name)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Couldn't find collection with that name!"})
		return
	}

	c.JSON(http.StatusOK, skinsFromCollection)
}

func (h *Handler) GetCollections(c *gin.Context) {
	collections, err := GetCollectionsJson(h.db)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Couldn't find collections!"})
		return
	}
	c.JSON(http.StatusOK, collections)
}

func (h *Handler) DropSkin(c *gin.Context) {
	collection := strings.TrimSpace(c.DefaultQuery("collection", ""))
	if collection == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Collection is required"})
		return
	}

	drop, err := DropSkinJson(h.db, collection)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "No skins found in this collection"})
		} else {
			// Log the actual error to console for debug
			fmt.Printf("DropSkinJson error: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't drop skin!"})
		}
		return
	}

	c.JSON(http.StatusOK, drop)
}
