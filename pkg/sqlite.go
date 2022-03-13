package pkg

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type DB struct {
	sql *sql.DB
}

func New() *DB {
	newDb, err := sql.Open("sqlite3", "./images.db")
	if err != nil {
		log.Fatalf("error to conn db %s", err)
	}
	return &DB{newDb}
}

type Image struct {
	Status string
	Id     int32
}

type Database interface {
	SaveToBase(filename string, image []byte) (Image, error)
	CheckBase(filename string) (Image, error)
}

func (database *DB) PrepareBase() error {
	statement, err := database.sql.Prepare("CREATE TABLE IF NOT EXISTS images (id INTEGER PRIMARY KEY, filename TEXT, image BLOB)")
	if err != nil {
		return err
	}

	statement.Exec()
	return nil
}

func (database *DB) SaveToBase(filename string, image []byte) (*Image, error) {
	statement, err := database.sql.Prepare("INSERT INTO images (filename, image) VALUES (?, ?)")
	if err != nil {
		return nil, fmt.Errorf("error from inserting in db: %s", err)
	}

	statement.Exec(filename, image)

	req, err := database.sql.Query("SELECT id FROM images WHERE filename = ?", filename)
	if err != nil {
		return nil, fmt.Errorf("error from saving to db: %s", err)
	}
	defer req.Close()

	insertedImage := Image{
		Status: "downloaded",
		Id:     -1,
	}

	for req.Next() {
		req.Scan(&insertedImage.Id)
	}
	if insertedImage.Id == -1 {
		return nil, fmt.Errorf("no rows with searching name %s", err)
	}

	return &insertedImage, nil
}

func (database *DB) CheckBase(filename string) (*Image, error) {
	rows, err := database.sql.Query("SELECT id FROM images WHERE filename = ?", filename)
	if err != nil {
		return nil, fmt.Errorf("error from checking db: %s", err)
	}
	defer rows.Close()

	response := Image{
		Status: "already downloaded",
		Id:     -1,
	}
	for rows.Next() {
		rows.Scan(&response.Id)
	}
	if response.Id == -1 {
		return nil, fmt.Errorf("no rows with searching name %s", err)
	}
	return &response, nil
}

func (database *DB) Clean(id int32) {
	database.sql.Exec("DELETE FROM images WHERE id = ?", id)
}

func (database *DB) Close() {
	database.sql.Close()
}
