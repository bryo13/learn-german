package main

import (
	"log"

	boltdb "github.com/BrianWork/learn-german-v2/boltDB"
)

func main() {
	db, err := boltdb.CreateDB()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
}
