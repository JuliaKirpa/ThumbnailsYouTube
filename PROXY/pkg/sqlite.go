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

func (database *DB) SaveToBase(filename string, image []byte) (*proto.Image, error) {
	statement, err := database.sql.Prepare("INSERT INTO images (filename, image) VALUES (?, ?)")
	if err != nil {
		return nil, err
	}

	statement.Exec(filename, image)

	req, err := database.sql.Query("SELECT id FROM images WHERE filename = " + filename)
	if err != nil {
		return nil, err
	}

	insertedImage := proto.Image{
		Status: "downloaded",
		Id:     0,
	}

	for req.Next() {
		req.Scan(&insertedImage.Id)
	}

	return &insertedImage, nil
}

func (database *DB) CheckBase(url string) (*proto.Image, error) {
	filename, _, err := ParceURL(url)
	if err != nil {
		return nil, err
	}

	image, err := database.sql.Query("SELECT status, id FROM images WHERE filename = " + filename)
	if err != nil {
		return nil, err
	}

	response := proto.Image{}
	for image.Next() {
		image.Scan(&response.Status, &response.Id)
	}
	return &response, nil
}
func (database *DB) Close() {
	database.sql.Close()
}
