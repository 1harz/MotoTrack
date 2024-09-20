package main

import (
	"gitlab.com/prof-caio-de-melo/user-service/internal/users"
	"gitlab.com/prof-caio-de-melo/user-service/internal/database"
	"log"
)

func main() {


	db := &database.MongoClient{}
	
	err := db.Start()

	if err != nil {
		log.Fatal("[FATAL] Could not connect with the database")
	}

	users.NewUserService(db.GetClient()).Start()

}
