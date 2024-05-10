package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type DBReadWriter struct {
	DbHost string
	DbPort int
	DbUser string
	DbPass string
	DbName string
}

type AppConfig struct {
	AppPort      int
	DbReadWriter DBReadWriter
}

func loadConfig() *AppConfig {
	var res = new(AppConfig)
	_, err := os.Stat(".env")
	if err == nil {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Failed to fetch .env file")
			return nil
		}
	}

	if value, found := os.LookupEnv("APP_PORT"); found {
		port, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal("Config : invalid server port", err.Error())
			return nil
		}
		res.AppPort = port
	} else {
		res.AppPort = 8080
	}

	if value, found := os.LookupEnv("DB_HOST"); found {
		res.DbReadWriter.DbHost = value
	}

	if value, found := os.LookupEnv("DB_PORT"); found {
		port, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal("Config : invalid db port", err.Error())
			return nil
		}
		res.DbReadWriter.DbPort = port
	}

	if value, found := os.LookupEnv("DB_USER"); found {
		res.DbReadWriter.DbUser = value
	}

	if value, found := os.LookupEnv("DB_PASS"); found {
		res.DbReadWriter.DbPass = value
	}

	if value, found := os.LookupEnv("DB_NAME"); found {
		res.DbReadWriter.DbName = value
	}
	return res
}

func GetConfig() *AppConfig {
	return loadConfig()
}
