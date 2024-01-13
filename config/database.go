package config

import (
	"database/sql"
	_ "github.com/sijms/go-ora/v2"
	goora "github.com/sijms/go-ora/v2"
	"github.com/spf13/viper"
)

// InitializeOracleDBCollection initializes OracleDB collection
func InitializeOracleDBCollection() *sql.DB {

	username := viper.GetString("ORACLE_DB_USERNAME")
	password := viper.GetString("ORACLE_DB_PASSWORD")

	connectString := viper.GetString("ORACLE_DB_CONNECTION_STRING")

	urlOptions := map[string]string{
		"TRACE FILE": "trace.log",
		"SSL VERIFY": "FALSE",
	}

	db, err := sql.Open("oracle", goora.BuildJDBC(username, password, connectString, urlOptions))

	if err != nil {
		panic(err.Error())
	}

	return db
}
