package boltdb

import (
	"fmt"
	"log"
	"time"

	"github.com/BrianWork/learn-german-v2/makedbdir"
	"github.com/boltdb/bolt"
)

// CreateDB func creates a bolt db returning an instance of bolt db
func CreateDB() (*bolt.DB, error) {

	// make db directory
	dirPath, _ := makedbdir.MakeDir()
	dbPath := fmt.Sprintf("%s/%s", dirPath, "learn-german-v2.db")

	fmt.Println(dbPath)
	db, err := bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 3 * time.Second})
	if err != nil {
		log.Println("Cant open db: ", err)
		return nil, err
	}

	fmt.Println("db created!")
	defer db.Close()

	// create bolt bucket
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("QandA"))
		if err != nil {
			log.Println("Cant open db: ", err)
			return err
		}
		return nil
	})

	if err != nil {
		log.Println("Cant open db: ", err)
		return nil, err
	}
	fmt.Println("bucket created!")

	return db, nil
}

// InsertIntoDB takes two strings and inserts them into the data returning an error if any
func InsertIntoDB(q, a string) error {
	db, err := CreateDB()
	if err != nil {
		return err
	}

	err = db.Update(func(tx *bolt.Tx) error {})
	if err != nil {
		log.Println("Cant insert into db: ", err)
		return err
	}
	return nil
}
