package dev

import (
	"log"

	"github.com/jexlor/cs2api/api"
	"github.com/jexlor/cs2api/db"
)

// this package is for development only, hide this handlers and remove endpoints in main.go file for production.

func DeleteSkinByNameJson(name string) error {
	_, err := db.DB.Exec(`DELETE FROM skins WHERE name = $1`, name)
	if err != nil {
		log.Printf("Error deleting skin: %v", err)
		return err
	}
	return nil
}

func UpdateSkinByNameJson(name string, updatedSkin api.Skin) error {
	query := `UPDATE skins SET 
		name = $1,
		rarity = $2,
		collection = $3,
		quality = $4,
		price = $5,
		url = $6
	WHERE name = $7`
	_, err := db.DB.Exec(query,
		updatedSkin.Name,
		updatedSkin.Rarity,
		updatedSkin.Collection,
		updatedSkin.Quality,
		updatedSkin.Price,
		updatedSkin.Url,
		name)
	if err != nil {
		log.Printf("Error updating skin: %v", err)
		return err
	}
	return nil
}
