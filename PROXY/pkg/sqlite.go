package pkg

import (
	"ThumbnailsYouTube_/PROXY/pkg/proto"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	sql *sql.DB
}

func ConnectToBase() (*DB, error) {
	newDb, err := sql.Open("sqlite3", "./images.db")
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

	req, err := database.sql.Query("SELECT id FROM images WHERE filename = ?", filename)
	if err != nil {
		return nil, fmt.Errorf("error from saving to db: ", err)
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

func (database *DB) CheckBase(filename string) (*proto.Image, error) {

	image, err := database.sql.Query("SELECT id FROM images WHERE filename = ?", filename)
	if err != nil {
		return nil, fmt.Errorf("error from checking db: ", err)
	}

	response := proto.Image{
		Status: "already downloaded",
		Id:     0,
	}
	for image.Next() {
		image.Scan(&response.Id)
	}
	return &response, nil
}

func (database *DB) Close() {
	database.sql.Close()
}
