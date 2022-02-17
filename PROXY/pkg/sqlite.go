package pkg

import (
	"ThumbnailsYouTube_/PROXY/pkg/proto"
	"database/sql"
)

type DB struct {
	sql *sql.DB
}

func ConnectToBase() (*DB, error) {
	newDb, err := sql.Open("sqlite3", "SQLite/images.db")
	if err != nil {
		return nil, err
	}
	statement, err := newDb.Prepare("CREATE TABLE IF NOT EXISTS images (id INTEGER PRIMARY KEY, filename TEXT, image BLOB)")
	if err != nil {
		return nil, err
	}
	statement.Exec()
	db := DB{sql: newDb}

	return &db, nil
}

func (database *DB) SaveToBase(filename, url string) (*proto.Image, error) {

	statement, err := database.sql.Prepare("INSERT INTO images (filename, image) VALUES (?, ?)")
	if err != nil {
		return nil, err
	}

	statement.Exec(filename)

	return &image, nil
}

func (database *DB) CheckBase(string string) (*proto.Image, error) {
	filename, url, err := ParceURL(string)
	if err != nil {
		return nil, err
	}

	image, err := database.sql.Query("SELECT status, id FROM images")
	if err != nil {
		database.SaveToBase(filename, url)
	}

	response := proto.Image{}
	for image.Next() {
		image.Scan(&response.Status, &response.Id)
	}
	return &response, nil
}
