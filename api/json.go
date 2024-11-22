package api

import (
	"fmt"
	"log"

	"github.com/jexlor/cs2api/db"
)

// structure of skins
type Skin struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Rarity     string `json:"rarity"`
	Collection string `json:"collection"`
	Quality    string `json:"quality"`
	Price      string `json:"price"` //type string in order to handle currency symbols
	Url        string `json:"url"`
}

func getAllSkinsJson() ([]Skin, error) {
	rows, err := db.DB.Query("SELECT * FROM skins")
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return nil, err
	}
	defer rows.Close()

	var skinsList []Skin

	for rows.Next() {
		var s Skin
		if err := rows.Scan(&s.Id, &s.Name, &s.Rarity, &s.Collection, &s.Quality, &s.Price, &s.Url); err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}
		skinsList = append(skinsList, s)
	}
	fmt.Println(len(skinsList))
	return skinsList, nil
}

func getSkinByIdJson(id int) (Skin, error) {
	var skin Skin
	err := db.DB.QueryRow(`SELECT * FROM skins WHERE id = $1`, id).Scan(
		&skin.Id,
		&skin.Name,
		&skin.Rarity,
		&skin.Collection,
		&skin.Quality,
		&skin.Price,
		&skin.Url,
	)
	if err != nil {
		return Skin{}, err
	}
	return skin, nil
}

func getSkinByNameJson(name string) (Skin, error) {
	var skin Skin
	err := db.DB.QueryRow(`SELECT * FROM skins WHERE name LIKE $1`, name).Scan(
		&skin.Id,
		&skin.Name,
		&skin.Rarity,
		&skin.Collection,
		&skin.Quality,
		&skin.Price,
		&skin.Url,
	)
	if err != nil {
		return Skin{}, err
	}
	return skin, nil
}

// func getCollectionByNameJson(name string) (Skin, error) {
// 	var collections Skin
// 	err := db.DB.Query(`SELECT * FROM skins WHERE name LIKE $1`, name).Scan(
// 		&skin.Id,
// 		&skin.Name,
// 		&skin.Rarity,
// 		&skin.Collection,
// 		&skin.Quality,
// 		&skin.Price,
// 		&skin.Url,
// 	)
// 	if err != nil {
// 		return Skin{}, err
// 	}
// 	return skin, nil
// }
