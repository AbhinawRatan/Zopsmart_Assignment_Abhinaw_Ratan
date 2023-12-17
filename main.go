package main

import (
	"fmt"
	"net/http"

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

// Initialize a slice of records
var records = []Record{
	{ID: 1, AlbumTitle: "The Dark Side of the Moon", Artist: "Pink Floyd", ReleaseYear: 1973, Genre: "Progressive Rock", Condition: "New", InStock: true},
	{ID: 2, AlbumTitle: "Abbey Road", Artist: "The Beatles", ReleaseYear: 1969, Genre: "Rock", Condition: "Good", InStock: true},
	// ... Add 8 more records
}

func main() {
	app := gofr.New()

	// Routes for records
	app.GET("/records", ListRecords)
	app.POST("/records", AddRecord)
	app.PUT("/records/:id/update", UpdateRecord)
	app.DELETE("/records/:id", RemoveRecord)

	// Start the server
	app.Start()
}

// ListRecords handler function
func ListRecords(ctx *gofr.Context) (interface{}, error) {
	return records, nil
}

// AddRecord handler function
func AddRecord(ctx *gofr.Context) (interface{}, error) {
	var newRecord Record
	if err := ctx.Bind(&newRecord); err != nil {
		return nil, fmt.Errorf("Invalid JSON format: %v", err)
	}

	// Assign a unique ID to the new record
	newRecord.ID = len(records) + 1

	// Add the new record to the slice
	records = append(records, newRecord)
	return newRecord, nil
}

// UpdateRecord handler function
func UpdateRecord(ctx *gofr.Context) (interface{}, error) {
	id := ctx.ParamAsInt("id")
	for i, record := range records {
		if record.ID == id {
			var updatedRecord Record
			if err := ctx.Bind(&updatedRecord); err != nil {
				return nil, fmt.Errorf("Invalid JSON format: %v", err)
			}
			records[i] = updatedRecord
			return updatedRecord, nil
		}
	}

	// If no record is found with the given ID
	return nil, gofr.NewHTTPError(http.StatusNotFound, "Record not found")
}

// RemoveRecord handler function
func RemoveRecord(ctx *gofr.Context) (interface{}, error) {
	id := ctx.ParamAsInt("id")
	for i, record := range records {
		if record.ID == id {
			records = append(records[:i], records[i+1:]...)
			return record, nil
		}
	}

	// If no record is found with the given ID
	return nil, gofr.NewHTTPError(http.StatusNotFound, "Record not found")
}
