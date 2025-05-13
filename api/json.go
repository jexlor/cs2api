package api

import (
	"database/sql"
	"fmt"

	"log"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"

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

func DropSkinJson(database *db.Database, collection string) (Skin, error) {
	rows, err := database.DB.Query(`
        SELECT id, name, weapon, rarity, collection, price, stattrack_price, url
        FROM skins
        WHERE collection = $1
    `, collection)
	if err != nil {
		return Skin{}, err
	}
	defer rows.Close()

	type entry struct {
		skin   Skin
		weight float64
	}
	var items []entry
	const alpha = 0.7 // chances (lower = more common, higher = rarer) after tests found out that 0.7 is the point of break.

	for rows.Next() {
		var s Skin
		if err := rows.Scan(
			&s.Id,
			&s.Name,
			&s.Weapon,
			&s.Rarity,
			&s.Collection,
			&s.Price,
			&s.StattrackPrice,
			&s.Url,
		); err != nil {
			return Skin{}, err
		}

		priceStr := strings.TrimPrefix(s.Price, "$")
		if strings.Contains(priceStr, "-") {
			priceStr = strings.Split(priceStr, "-")[0] // take the lower bound
		}
		priceStr = strings.ReplaceAll(priceStr, ",", "") // remove commas
		priceStr = strings.TrimSpace(priceStr)

		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil || price <= 0 {
			return Skin{}, fmt.Errorf("invalid skin price: %s", s.Price)
		}

		weight := math.Pow(1.0/price, alpha)
		items = append(items, entry{skin: s, weight: weight})
	}

	if len(items) == 0 {
		return Skin{}, sql.ErrNoRows
	}

	var totalWeight float64
	for _, item := range items {
		totalWeight += item.weight
	}

	rand.Seed(time.Now().UnixNano())
	target := rand.Float64() * totalWeight

	var cumulative float64
	for _, item := range items {
		cumulative += item.weight
		if target <= cumulative {
			return item.skin, nil
		}
	}

	return items[len(items)-1].skin, nil
}
