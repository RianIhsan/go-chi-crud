package db

import (
	"database/sql"
	"fmt"
	"github.com/RianIhsan/go-chi-crud/config"
	_ "github.com/lib/pq"
	"log"
)

func SetupDatabase() (*sql.DB, error) {
	appConfig := config.GetConfig()
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		appConfig.DbReadWriter.DbHost,
		appConfig.DbReadWriter.DbPort,
		appConfig.DbReadWriter.DbUser,
		appConfig.DbReadWriter.DbPass,
		appConfig.DbReadWriter.DbName)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}
