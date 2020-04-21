package demo

import (
	"github.com/boltdb/bolt"
	"log"
)

func OpenDB() (*bolt.DB, error) {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}

func CloseDB(db *bolt.DB) {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
}
