package helpers

import (
	"fmt"
	"os"
)

func CheckEnvVariables() {
	var dbHost = os.Getenv("DB_HOST")
	// var dbPort = os.Getenv("DB_PORT")
	var dbName = os.Getenv("DB_NAME")
	var dbUser = os.Getenv("DB_USER")
	var dbPassword = os.Getenv("DB_PASSWORD")

	if len(dbHost) == 0 || len(dbName) == 0 || len(dbUser) == 0 || len(dbPassword) == 0 {
		panic("Please set database variables: DB_HOST, DB_NAME, DB_USER, DB_PASSWORD")
	}

	var jwtSecret = os.Getenv("JWT_SECRET")
	if len(jwtSecret) == 0 {
		var randStr = RandStringBytes(20)
		fmt.Printf("JWT_SECRET not set, use generated string instead: %s\n", randStr)
		os.Setenv("JWT_SECRET", randStr)
	}

}
