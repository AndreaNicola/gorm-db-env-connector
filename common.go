package gorm_db_env_connector

import "os"

func resolveEnvOrDefault(envVar string, defaultValue string) string {

	res := os.Getenv(envVar)
	if res == "" {
		res = defaultValue
	}
	return res

}
