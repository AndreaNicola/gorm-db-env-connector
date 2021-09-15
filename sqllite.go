package gorm_db_env_connector

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type SqlliteParams struct {
	DatabaseName string
}

type SqlliteEnv struct {
	DatabaseNameEnvVar  string
	DatabaseNameDefault string
}

func SqlliteConnect(params SqlliteParams) *gorm.DB {

	log.Println("Initializing connection to database " + params.DatabaseName)
	db, err := gorm.Open(sqlite.Open(params.DatabaseName), &gorm.Config{})
	if err != nil {
		log.Fatal("Connection to database " + params.DatabaseName + " failed")
	}

	log.Println("Connection to database " + params.DatabaseName + " completed")
	return db

}
