package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"fmt"

	"github.com/stretchr/testify/assert"
	"gofr.dev/pkg/gofr"
)

// Record struct to represent a vinyl record
type Record struct {
	ID          int    `json:"id"`
	AlbumTitle  string `json:"albumTitle"`
	Artist      string `json:"artist"`
	ReleaseYear int    `json:"releaseYear"`
	Genre       string `json:"genre"`
	Condition   string `json:"condition"` // e.g., "New", "Good", "Worn"
	InStock     bool   `json:"inStock"`
}

func setupApp() *gofr.App {
	app := gofr.New()

	// Routes for records
	app.GET("/records", ListRecords)
	app.POST("/records", AddRecord)
	app.PUT("/records/:id/update", UpdateRecord)
	app.DELETE("/records/:id", RemoveRecord)

	return app
}

func TestRecordCRUD(t *testing.T) {
	// Setup
	app := setupApp()

	// Test AddRecord
	newRecord := Record{
		AlbumTitle:  "The Dark Side of the Moon",
		Artist:      "Pink Floyd",
		ReleaseYear: 1973,
		Genre:       "Progressive Rock",
		Condition:   "New",
		InStock:     true,
	}
	newRecordJSON, _ := json.Marshal(newRecord)
	reqAddRecord := httptest.NewRequest("POST", "/records", bytes.NewBuffer(newRecordJSON))
	reqAddRecord.Header.Set("Content-Type", "application/json")
	respAddRecord := httptest.NewRecorder()
	app.ServeHTTP(respAddRecord, reqAddRecord)

	assert.Equal(t, http.StatusOK, respAddRecord.Code)

	// Test ListRecords
	reqListRecords := httptest.NewRequest("GET", "/records", nil)
	respListRecords := httptest.NewRecorder()
	app.ServeHTTP(respListRecords, reqListRecords)

	assert.Equal(t, http.StatusOK, respListRecords.Code)

	var retrievedRecords []Record
	err := json.Unmarshal(respListRecords.Body.Bytes(), &retrievedRecords)
	assert.NoError(t, err)
	assert.NotEmpty(t, retrievedRecords)

	// Test UpdateRecord - Assuming updating the stock status
	idToUpdate := retrievedRecords[0].ID
	updatedRecord := retrievedRecords[0]
	updatedRecord.InStock = false
	updateRecordJSON, _ := json.Marshal(updatedRecord)
	reqUpdateRecord := httptest.NewRequest("PUT", fmt.Sprintf("/records/%d/update", idToUpdate), bytes.NewBuffer(updateRecordJSON))
	reqUpdateRecord.Header.Set("Content-Type", "application/json")
	respUpdateRecord := httptest.NewRecorder()
	app.ServeHTTP(respUpdateRecord, reqUpdateRecord)

	assert.Equal(t, http.StatusOK, respUpdateRecord.Code)

	// Test RemoveRecord
	idToRemove := retrievedRecords[0].ID
	reqRemoveRecord := httptest.NewRequest("DELETE", fmt.Sprintf("/records/%d", idToRemove), nil)
	respRemoveRecord := httptest.NewRecorder()
	app.ServeHTTP(respRemoveRecord, reqRemoveRecord)

	assert.Equal(t, http.StatusOK, respRemoveRecord.Code)

	// Verify ListRecords after removal
	reqListRecordsAfterRemove := httptest.NewRequest("GET", "/records", nil)
	respListRecordsAfterRemove := httptest.NewRecorder()
	app.ServeHTTP(respListRecordsAfterRemove, reqListRecordsAfterRemove)

	assert.Equal(t, http.StatusOK, respListRecordsAfterRemove.Code)

	var recordsAfterRemove []Record
	err = json.Unmarshal(respListRecordsAfterRemove.Body.Bytes(), &recordsAfterRemove)
	assert.NoError(t, err)
	assert.True(t, len(recordsAfterRemove) < len(retrievedRecords))
}
