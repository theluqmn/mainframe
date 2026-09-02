package utils

import (
	"fmt"
	"time"
	"log"
	"database/sql"
	_ "modernc.org/sqlite"
)

func NewDB(path string) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s?_pragma=journal_mode(WAL)&_pragma=busy_timeout(5000)", path)

	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(time.Hour)

	err = db.Ping();
	if err != nil {
		return nil, err
	}

	// create admins table
	query := "CREATE TABLE IF NOT EXISTS admins (id TEXT PRIMARY KEY, password TEXT, salt TEXT)"
	_, err = db.Exec(query)
	if err != nil {
		log.Fatalf("error creating table: %v", err)
	}

	// insert initial admin user
	stmt, err := db.Prepare("INSERT INTO admins (id, password, salt) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatalf("error preparing statement: %v", err)
	}
	defer stmt.Close()
	
	_, err = stmt.Exec("admin", "123456", "im flying intercontinental with you")
	if err != nil {
		log.Fatalf("error executing statement: %v", err)
	}

	return db, nil
}
