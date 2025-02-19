package api

import (
	"fmt"
	"log"

	"github.com/jexlor/cs2api/db"
)

func GetAllSkinsJson(database *db.Database) ([]Skin, error) {
	rows, err := database.DB.Query("SELECT * FROM skins")
	if err != nil {
		log.Printf("Error executing query: %v", err) //remove raw database error messages for production
		return nil, err
	}

	defer rows.Close()

	var skinsList []Skin

	for rows.Next() {
		var s Skin
		if err := rows.Scan(&s.Id, &s.Name, &s.Weapon, &s.Rarity, &s.Collection, &s.Price, &s.StattrackPrice, &s.Url); err != nil {
			log.Printf("Error scanning row: %v", err) //remove raw database error messages for production
			return nil, err
		}
		skinsList = append(skinsList, s)
	}
	fmt.Println(len(skinsList))
	return skinsList, nil
}

func GetSkinByIdJson(database *db.Database, id int) (Skin, error) {
	var skin Skin
	err := database.DB.QueryRow(`SELECT * FROM skins WHERE id = $1`, id).Scan(
		&skin.Id,
		&skin.Name,
		&skin.Weapon,
		&skin.Rarity,
		&skin.Collection,
		&skin.Price,
		&skin.StattrackPrice,
		&skin.Url,
	)
	if err != nil {
		return Skin{}, err
	}
	return skin, nil
}

func GetSkinByNameJson(database *db.Database, name string) (Skin, error) {
	var skin Skin
	err := database.DB.QueryRow(`SELECT * FROM skins WHERE name = $1`, name).Scan(
		&skin.Id,
		&skin.Name,
		&skin.Weapon,
		&skin.Rarity,
		&skin.Collection,
		&skin.Price,
		&skin.StattrackPrice,
		&skin.Url,
	)
	if err != nil {
		return Skin{}, err
	}
	return skin, nil
}

func GetCollectionByNameJson(database *db.Database, name string) ([]Skin, error) {
	rows, err := database.DB.Query("SELECT * FROM skins WHERE collection = $1", name)
	if err != nil {
		log.Printf("Error executing query: %v", err) //remove raw database error messages for production
		return nil, err
	}

	defer rows.Close()

	var skinsFromCollection []Skin
	for rows.Next() {
		var s Skin
		if err := rows.Scan(&s.Id, &s.Name, &s.Weapon, &s.Rarity, &s.Collection, &s.Price, &s.StattrackPrice, &s.Url); err != nil {
			log.Printf("Error scanning row: %v", err) //remove raw database error messages for production
			return nil, err
		}
		skinsFromCollection = append(skinsFromCollection, s)
	}
	return skinsFromCollection, nil
}
func GetCollectionsJson(database *db.Database) ([]Col, error) {
	rows, err := database.DB.Query(`SELECT DISTINCT collection FROM skins`)
	if err != nil {
		log.Printf("Error executing query: %v", err) //remove raw database error messages for production
		return nil, err
	}

	defer rows.Close()

	var collections []Col

	for rows.Next() {
		var c Col
		if err := rows.Scan(&c.Collection); err != nil {
			log.Printf("Error scanning row: %v", err) //remove raw database error messages for production
			return nil, err
		}
		collections = append(collections, c)
	}
	return collections, nil
}
