package dev

import (
	"log"

	"github.com/jexlor/cs2api/api"
	"github.com/jexlor/cs2api/db"
)

func DeleteSkinByNameJson(database *db.Database, name string) error {
	_, err := database.DB.Exec(`DELETE FROM skins WHERE name = $1`, name)
	if err != nil {
		log.Printf("Error deleting skin: %v", err)
		return err
	}
	return nil
}

func UpdateSkinByNameJson(database *db.Database, name string, updatedSkin api.Skin) error {
	query := `UPDATE skins SET 
		name = $1,
		weapon = $2,
		rarity = $3,
		collection = $4,
		price = $5,
		stattrack_price = $6,
		url = $7
	WHERE name = $8`
	_, err := database.DB.Exec(query,
		updatedSkin.Name,
		updatedSkin.Weapon,
		updatedSkin.Rarity,
		updatedSkin.Collection,
		updatedSkin.Price,
		updatedSkin.StattrackPrice,
		updatedSkin.Url,
		name)
	if err != nil {
		log.Printf("Error updating skin: %v", err)
		return err
	}
	return nil
}
