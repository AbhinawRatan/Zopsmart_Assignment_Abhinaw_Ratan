package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type Record struct {
	ID          int    `json:"id"`
	AlbumTitle  string `json:"albumTitle"`
	Artist      string `json:"artist"`
	ReleaseYear int    `json:"releaseYear"`
	Genre       string `json:"genre"`
	Condition   string `json:"condition"`
	InStock     bool   `json:"inStock"`
}

func main() {
	var err error
	db, err = sql.Open("sqlite3", "vinyl_records.db")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	createTable()
	http.HandleFunc("/", operations)
	log.Fatal(http.ListenAndServe(":4000", nil))
}

func createTable() {
	createSQL := `CREATE TABLE IF NOT EXISTS records (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		albumTitle TEXT,
		artist TEXT,
		releaseYear INTEGER,
		genre TEXT,
		condition TEXT,
		inStock BOOLEAN
	);`
	if _, err := db.Exec(createSQL); err != nil {
		log.Fatalf("Error creating table: %v", err)
	}
}

func operations(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "<h1>Hi,The Api is Live </h1>")
	case "/delete":
		deleteRecord(w, r)
	case "/add":
		addRecord(w, r)
	case "/update":
		updateRecord(w, r)
	case "/view":
		viewRecords(w, r)
	default:
		http.Error(w, "404 not found.", http.StatusNotFound)
	}
}

func addRecord(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	albumTitle := r.FormValue("albumTitle")
	artist := r.FormValue("artist")
	releaseYear, err := strconv.Atoi(r.FormValue("releaseYear"))
	if err != nil {
		http.Error(w, "Invalid release year.", http.StatusBadRequest)
		return
	}
	genre := r.FormValue("genre")
	condition := r.FormValue("condition")
	inStock, err := strconv.ParseBool(r.FormValue("inStock"))
	if err != nil {
		http.Error(w, "Invalid in stock value.", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("INSERT INTO records (albumTitle, artist, releaseYear, genre, condition, inStock) VALUES (?, ?, ?, ?, ?, ?)", albumTitle, artist, releaseYear, genre, condition, inStock)
	if err != nil {
		http.Error(w, "Error adding record.", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Record added successfully")
}

func viewRecords(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	rows, err := db.Query("SELECT * FROM records")
	if err != nil {
		http.Error(w, "Error reading records.", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var records []Record
	for rows.Next() {
		var rec Record
		if err := rows.Scan(&rec.ID, &rec.AlbumTitle, &rec.Artist, &rec.ReleaseYear, &rec.Genre, &rec.Condition, &rec.InStock); err != nil {
			http.Error(w, "Error scanning records.", http.StatusInternalServerError)
			return
		}
		records = append(records, rec)
	}

	for _, rec := range records {
		fmt.Fprintf(w, "ID: %d, AlbumTitle: %s, Artist: %s, ReleaseYear: %d, Genre: %s, Condition: %s, InStock: %t\n", rec.ID, rec.AlbumTitle, rec.Artist, rec.ReleaseYear, rec.Genre, rec.Condition, rec.InStock)
	}
}

func updateRecord(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID.", http.StatusBadRequest)
		return
	}

	albumTitle := r.FormValue("albumTitle")
	artist := r.FormValue("artist")
	releaseYear, err := strconv.Atoi(r.FormValue("releaseYear"))
	if err != nil {
		http.Error(w, "Invalid release year.", http.StatusBadRequest)
		return
	}
	genre := r.FormValue("genre")
	condition := r.FormValue("condition")
	inStock, err := strconv.ParseBool(r.FormValue("inStock"))
	if err != nil {
		http.Error(w, "Invalid in stock value.", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("UPDATE records SET albumTitle=?, artist=?, releaseYear=?, genre=?, condition=?, inStock=? WHERE id=?", albumTitle, artist, releaseYear, genre, condition, inStock, id)
	if err != nil {
		http.Error(w, "Error updating record.", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Record updated successfully")
}

func deleteRecord(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID.", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("DELETE FROM records WHERE id=?", id)
	if err != nil {
		http.Error(w, "Error deleting record.", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Record deleted successfully")
}
