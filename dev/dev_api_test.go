// package dev

// import (
// 	"log"
// 	"testing"

// 	"github.com/DATA-DOG/go-sqlmock"
// 	"github.com/jexlor/cs2api/api"
// 	"github.com/jexlor/cs2api/db"
// 	"github.com/stretchr/testify/assert"
// )

// // Setup a mock database connection
// func setupMockDB() (sqlmock.Sqlmock, func()) {
// 	// Create a mock DB connection
// 	mockDB, mock, err := sqlmock.New()
// 	if err != nil {
// 		log.Fatalf("Error setting up mock DB: %v", err)
// 	}

// 	// Replace the real DB with the mock DB for testing
// 	db.DB = mockDB // Override the real DB with mock DB only during tests

// 	// Return the mock object and a cleanup function
// 	return mock, func() {
// 		// Close the mock DB when done
// 		mockDB.Close()
// 	}
// }

// // Test the DeleteSkinByNameJson function
// func TestDeleteSkinByNameJson(t *testing.T) {
// 	mock, cleanup := setupMockDB()
// 	defer cleanup()

// 	// Expect the DELETE query to be called once with the skin name "TestSkin"
// 	mock.ExpectExec("DELETE FROM skins WHERE name = $1").
// 		WithArgs("TestSkin").
// 		WillReturnResult(sqlmock.NewResult(1, 1))

// 	// Call the function to test
// 	err := DeleteSkinByNameJson("TestSkin")

// 	// Assert that no error occurred
// 	assert.NoError(t, err)

// 	// Assert that the expected query was executed
// 	err = mock.ExpectationsWereMet()
// 	assert.NoError(t, err)
// }

// // Test the UpdateSkinByNameJson function
// func TestUpdateSkinByNameJson(t *testing.T) {
// 	mock, cleanup := setupMockDB()
// 	defer cleanup()

// 	// Define the updated skin data
// 	updatedSkin := api.Skin{
// 		Name:       "UpdatedSkin",
// 		Rarity:     "Rare",
// 		Collection: "New Collection",
// 		Quality:    "Minimal Wear",
// 		Price:      "$50.0",
// 		Url:        "http://new-url.com",
// 	}

// 	// Expect the UPDATE query to be called with the updated skin details
// 	mock.ExpectExec(`UPDATE skins SET`).
// 		WithArgs(
// 			updatedSkin.Name,
// 			updatedSkin.Rarity,
// 			updatedSkin.Collection,
// 			updatedSkin.Quality,
// 			updatedSkin.Price,
// 			updatedSkin.Url,
// 			"OldSkinName",
// 		).
// 		WillReturnResult(sqlmock.NewResult(1, 1))

// 	// Call the function to test
// 	err := UpdateSkinByNameJson("OldSkinName", updatedSkin)

// 	// Assert that no error occurred
// 	assert.NoError(t, err)

// 	// Assert that the expected query was executed
// 	err = mock.ExpectationsWereMet()
// 	assert.NoError(t, err)
// }

package dev

import "testing"

func TestExample(t *testing.T) {
	t.Log("This is a placeholder test.")
}
