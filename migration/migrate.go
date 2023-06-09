package migration

import (
	authorData "github.com/iqbalpradipta/HTTP-RESTful-API-using-Go-and-a-SQL-Database/feature/author/data"
	bookData "github.com/iqbalpradipta/HTTP-RESTful-API-using-Go-and-a-SQL-Database/feature/book/data"
	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB)  {
	db.AutoMigrate(authorData.Author{})
	db.AutoMigrate(bookData.Book{})
}