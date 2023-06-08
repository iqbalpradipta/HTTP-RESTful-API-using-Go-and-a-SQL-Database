package config

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	SERVER_PORT		int
	DB_DRIVER		string
	DB_USERNAME		string
	DB_PASSWORD		string
	DB_HOST			string
	DB_PORT			int
	DB_NAME			string
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = InitConfig()
	}
	return appConfig
}

func InitConfig() *AppConfig {
	var defaultConfig AppConfig

	err := godotenv.Load(".env")
  	if err != nil {
    log.Fatal("Error loading .env file")
  	}

	serverPortConv, errConv := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if errConv != nil {
		log.Fatal("error parse SERVER PORT")
		return nil
	}

	defaultConfig.SERVER_PORT = serverPortConv
	defaultConfig.DB_DRIVER = os.Getenv("DB_DRIVER")
	defaultConfig.DB_USERNAME = os.Getenv("DB_USERNAME")	
	defaultConfig.DB_HOST = os.Getenv("DB_HOST")
	defaultConfig.DB_NAME = os.Getenv("DB_NAME")
	defaultConfig.DB_PASSWORD = os.Getenv("DB_PASSWORD")

	port, errPortConv := strconv.Atoi(os.Getenv("DB_PORT"))
	if errPortConv != nil {
		log.Fatal("error parse DB PORT")
		return nil
	}

	defaultConfig.DB_PORT = port
	return &defaultConfig
}