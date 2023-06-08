package migration

import (
	authorData "github.com/iqbalpradipta/HTTP-RESTful-API-using-Go-and-a-SQL-Database/feature/author/data"
	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB)  {
	db.AutoMigrate(authorData.Author{})
}