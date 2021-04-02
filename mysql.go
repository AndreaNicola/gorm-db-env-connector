package gorm_db_env_connector

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func MySQLConnect(dbUrlEnv string, dbSchemaEnv string, dbUsernameEnv string, dbPasswordEnv string) *gorm.DB {

	dbUrl := os.Getenv(dbUrlEnv)
	dbSchema := os.Getenv(dbSchemaEnv)
	dbUsername := os.Getenv(dbUsernameEnv)
	dbPassword := os.Getenv(dbPasswordEnv)

	log.Debug("Initializing connection to database " + dbUrl)

	dsn := dbUsername + ":" + dbPassword + "@tcp(" + dbUrl + ")/" + dbSchema + "?parseTime=true"

	_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Connection to database " + dbUrl + " failed")
		panic("Connection to database failed")
	}

	log.Debug("Connection to database " + dbUrl + " completed")

	return _db
}
