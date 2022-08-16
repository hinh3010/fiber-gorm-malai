package migration

import (
	"log"
	"malai/database"
	"malai/model/emtity"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&emtity.UserEmtity{})
	if err != nil {
		log.Panicln(err)
	}
	log.Println("database migration successful")
}
