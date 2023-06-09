package data

import (
	"github.com/iqbalpradipta/HTTP-RESTful-API-using-Go-and-a-SQL-Database/feature/user"
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name    string
	Country string
	Books   []Book `gorm:"many2many:author_books;"`
}

type Book struct {
	gorm.Model
	Title           string
	PublicationYear int
	ISBN            string
	Authors         []Author `gorm:"many2many:author_books;"`
}

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}


func fromCore(data user.UserCore) User {
	return User{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}
}

func toCore(data User) user.UserCore {
	return user.UserCore{
		ID:       int(data.ID),
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}
}