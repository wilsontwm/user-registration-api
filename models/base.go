package models

import (
	"github.com/joho/godotenv"
	"github.com/wilsontwm/user-registration"
	"os"
	"strconv"
)

var IsActivationRequired = false

// Initialization
func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	IsActivationRequired, _ = strconv.ParseBool(os.Getenv("is_activation_required"))

	// Start of initialization of the user registration module
	dbConfig := userreg.DBConfig{
		Driver:                 os.Getenv("db_type"),
		Username:               os.Getenv("db_user"),
		Password:               os.Getenv("db_pass"),
		Host:                   os.Getenv("db_host"),
		DBName:                 os.Getenv("db_name"),
		InstanceConnectionName: os.Getenv("db_instance"),
	}

	tableName := "users"
	userreg.Initialize(dbConfig)
	userreg.Config(userreg.TableName(tableName), userreg.UserActivation(IsActivationRequired), userreg.MigrateDatabase())
	// End of initialization of the user registration module
}
