package models

import (
	"github.com/joho/godotenv"
	"github.com/wilsontwm/user-registration"
	"os"
)

// Initialization
func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// Start of initialization of the user registration module
	dbConfig := userreg.DBConfig{
		Driver:   os.Getenv("db_type"),
		Username: os.Getenv("db_user"),
		Password: os.Getenv("db_pass"),
		Host:     os.Getenv("db_host"),
		DBName:   os.Getenv("db_name"),
	}

	tableName := "tests"
	userreg.Initialize(dbConfig)
	userreg.Config(userreg.TableName(tableName))
	// End of initialization of the user registration module
}
