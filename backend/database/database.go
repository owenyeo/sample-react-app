package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/owenyeo/sample-react-app/backend/models"
)

type Database struct {
	db *sql.DB
}

func NewDatabase(dataSourceName string) (*Database, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to the database!")

	return &Database{db: db}, nil
}

func (d *Database) Close() error {
	if d.db == nil {
		return nil
	}
	return d.db.Close()
}

func GetDB() (*Database, error) {
	return NewDatabase("user=owenyeoo password=password dbname=cvwo sslmode=disable")
}

func (d *Database) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return d.db.Query(query, args...)
}

func (d *Database) Exec(query string, args ...interface{}) (sql.Result, error) {
	return d.db.Exec(query, args...)
}

func (d *Database) AddUser(newUser models.User) error {
	_, err := d.db.Exec("INSERT INTO users (name) VALUES ($1)", newUser.Name)

	if err != nil {
		return err
	}

	return nil
}

func (d *Database) UserExists(username string) (bool, error) {
	var exists bool
	err := d.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE name = $1)", username).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil // User does not exist
		}
		return false, err // Some other error occurred
	}
	return exists, nil // User exists or error while checking
}


