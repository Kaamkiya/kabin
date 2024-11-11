package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Paste struct {
	ID string `json:"id"`
	Content string `json:"content"`
}

var db *sql.DB

func Open() (err error) {
	db, err = sql.Open("sqlite3", "./pastes.db")
	return err
}

func Init() error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS pastes (
			id TEXT PRIMARY KEY,
			content TEXT
	);`)
	return err
}

func GetPaste(id string) (Paste, error) {
	row := db.QueryRow("SELECT * FROM pastes WHERE id = ?", id)
	
	p := Paste{}
	err := row.Scan(&p.ID, &p.Content)

	return p, err
}

func AddPaste(p Paste) error {
	_, err := db.Exec("INSERT INTO pastes (id, content) VALUES (?, ?);", p.ID, p.Content)

	return err
}

func Close() error {
	return db.Close()
}
