package postgresql

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/iqbalpradipta/HTTP-RESTful-API-using-Go-and-a-SQL-Database/config"
)
func InitDB(cfg *config.AppConfig) *gorm.DB{
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	cfg.DB_HOST,
	cfg.DB_PORT,
	cfg.DB_USERNAME,
	cfg.DB_PASSWORD,
	cfg.DB_NAME,
	)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("error connect to DB", err)
	}

	return db
}