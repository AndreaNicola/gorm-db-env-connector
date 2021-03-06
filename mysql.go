package gorm_db_env_connector

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type MySqlParams struct {
	DbUrl      string
	DbSchema   string
	DbUsername string
	DbPassword string
}

type MySqlEnv struct {
	DbUrlEnvVar       string
	DbSchemaEnvVar    string
	DbUsernameEnvVar  string
	DbPasswordEnvVar  string
	DbUrlDefault      string
	DbSchemaDefault   string
	DbUsernameDefault string
	DbPasswordDefault string
}

func (env *MySqlEnv) resolve() MySqlParams {

	return MySqlParams{
		DbUrl:      resolveEnvOrDefault(env.DbUrlEnvVar, env.DbUrlDefault),
		DbSchema:   resolveEnvOrDefault(env.DbSchemaEnvVar, env.DbSchemaDefault),
		DbUsername: resolveEnvOrDefault(env.DbUsernameEnvVar, env.DbUsernameDefault),
		DbPassword: resolveEnvOrDefault(env.DbPasswordEnvVar, env.DbPasswordDefault),
	}
	
}

func MySQLConnect(mysqlParams MySqlParams) *gorm.DB {

	databaseName :=  mysqlParams.DbUrl + "/"+ mysqlParams.DbSchema

	log.Println("Initializing connection to database " + databaseName)

	dsn := mysqlParams.DbUsername + ":" + mysqlParams.DbPassword + "@tcp(" + mysqlParams.DbUrl + ")/" + mysqlParams.DbSchema + "?parseTime=true"

	_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Connection to database " + databaseName + " failed")
	}

	log.Println("Connection to database " + databaseName + " completed")

	_ddb, err := _db.DB()

	if err != nil {
		panic("Setting connection pool failed")
	}

	_ddb.SetMaxIdleConns(10)
	_ddb.SetMaxOpenConns(50)

	return _db
}

func MySQLConnectEnv(mysqlEnv MySqlEnv) *gorm.DB {
	return MySQLConnect(mysqlEnv.resolve())
}
