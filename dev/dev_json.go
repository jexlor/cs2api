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
		weapon = $2,
		rarity = $3,
		collection = $4,
		quality = $5,
		price = $6,
		statrack_price = $7,
		url = $8
	WHERE name = $9`
	_, err := db.DB.Exec(query,
		updatedSkin.Name,
		updatedSkin.Weapon,
		updatedSkin.Rarity,
		updatedSkin.Collection,
		updatedSkin.Price,
		updatedSkin.Url,
		name)
	if err != nil {
		log.Printf("Error updating skin: %v", err)
		return err
	}
	return nil
}
