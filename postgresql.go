package gorm_db_env_connector

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type PgParams struct {
	Url      string
	Database string
	Username string
	Password string
	Port     string
	SSLMode  string
}

type PgEnv struct {
	UrlEnvVar       string
	PortEnvVar      string
	DatabaseEnvVar  string
	UsernameEnvVar  string
	PasswordEnvVar  string
	SSLModeEnvVar   string
	UrlDefault      string
	PortDefault     string
	DatabaseDefault string
	UsernameDefault string
	PasswordDefault string
	SSLModeDefault  string
}

func (env *PgEnv) resolve() PgParams {

	return PgParams{
		Url:      resolveEnvOrDefault(env.UrlEnvVar, env.UrlDefault),
		Database: resolveEnvOrDefault(env.DatabaseEnvVar, env.DatabaseDefault),
		Username: resolveEnvOrDefault(env.UsernameEnvVar, env.UsernameDefault),
		Password: resolveEnvOrDefault(env.PasswordEnvVar, env.PasswordDefault),
		Port:     resolveEnvOrDefault(env.PortEnvVar, env.PortDefault),
		SSLMode:  resolveEnvOrDefault(env.SSLModeEnvVar, env.SSLModeDefault),
	}

}

func PgConnectEnv(env PgEnv) *gorm.DB {
	return PgConnect(env.resolve())
}

func PgConnect(pgParams PgParams) *gorm.DB {

	log.Println("Initializing connection to database " + pgParams.Url)
	dsn := "host=" + pgParams.Url + " user=" + pgParams.Username + " password=" + pgParams.Password + " dbname=" + pgParams.Database + " port=" + pgParams.Port + " sslmode=" + pgParams.SSLMode

	_db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Connection to database " + pgParams.Url + " failed")
	}

	log.Println("Connection to database " + pgParams.Url + " completed")

	_ddb, err := _db.DB()

	if err != nil {
		panic("Setting connection pool failed")
	}

	_ddb.SetMaxIdleConns(10)
	_ddb.SetMaxOpenConns(50)

	return _db

}
