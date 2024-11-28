package main

import (
	"database/sql"
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jexlor/cs2api/api"
	"github.com/stretchr/testify/assert"
)

func waitForDB(db *sql.DB) {
	for {
		err := db.Ping()
		if err == nil {
			log.Println("Database initialized successfully.")
			break
		}
		log.Println("Waiting for database to initialize...")
		time.Sleep(2 * time.Second)
	}
}

func TestGetSkinByIdJson(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error occurred when opening a mock DB: %s", err)
	}
	defer db.Close()

	waitForDB(db)

	mock.ExpectQuery(`SELECT \* FROM skins WHERE id = \$1`).WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "rarity", "collection", "quality", "price", "url"}).
			AddRow(1, "Skin A", "Rare", "Collection A", "High", "$100", "http://example.com"))

	var skin api.Skin
	err = db.QueryRow("SELECT * FROM skins WHERE id = $1", 1).Scan(&skin.Id, &skin.Name, &skin.Rarity, &skin.Collection, &skin.Quality, &skin.Price, &skin.Url)
	if err != nil {
		t.Fatalf("Error retrieving skin: %s", err)
	}

	assert.Equal(t, 1, skin.Id)
	assert.Equal(t, "Skin A", skin.Name)
	assert.Equal(t, "Rare", skin.Rarity)
	assert.Equal(t, "Collection A", skin.Collection)
	assert.Equal(t, "High", skin.Quality)
	assert.Equal(t, "$100", skin.Price)
	assert.Equal(t, "http://example.com", skin.Url)

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("there were unmet expectations: %s", err)
	}
}
