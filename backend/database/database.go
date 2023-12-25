package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq" // PostgreSQL driver
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

	sqlFile, err := os.ReadFile("backend/database/cvwo.sql")
	if err != nil {
		fmt.Println("Error reading the schema file:", err)
	}

	// Execute the schema
	_, err = db.Exec(string(sqlFile))
	if err != nil {
		fmt.Println("Error executing the schema:", err)
	}

	fmt.Println("Schema executed successfully!")

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

func (d *Database) QueryRow(query string, args ...interface{}) *sql.Row {
	return d.db.QueryRow(query, args...)
}
