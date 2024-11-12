package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Paste struct {
	ID string `json:"id"`
	Content string `json:"content"`
	Language string `json:"language"`
}

var db *sql.DB

func Open() (err error) {
	db, err = sql.Open("sqlite3", "./pastes.db")
	return err
}

func Init() error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS pastes (
			id TEXT PRIMARY KEY,
			content TEXT,
			language TEXT
	);`)
	return err
}

func GetPaste(id string) (Paste, error) {
	row := db.QueryRow("SELECT * FROM pastes WHERE id = ?", id)
	
	p := Paste{}
	err := row.Scan(&p.ID, &p.Content, &p.Language)

	return p, err
}

func AddPaste(p Paste) error {
	_, err := db.Exec("INSERT INTO pastes (id, content, language) VALUES (?, ?, ?);", p.ID, p.Content, p.Language)

	return err
}

func Close() error {
	return db.Close()
}
