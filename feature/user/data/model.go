package data

import (
	dataAuthor "github.com/iqbalpradipta/HTTP-RESTful-API-using-Go-and-a-SQL-Database/feature/author/data"
	dataBook "github.com/iqbalpradipta/HTTP-RESTful-API-using-Go-and-a-SQL-Database/feature/book/data"
	"github.com/iqbalpradipta/HTTP-RESTful-API-using-Go-and-a-SQL-Database/feature/user"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name		string 
	Email		string `gorm:"unique"`
	Password	string
	BookId 		int
	AuthorId 	int
	Book		[]dataBook.Book		`gorm:"foreignkey:BookID"`
	Author		[]dataAuthor.Author	`gorm:"foreignkey:AuthorID"`
}

func fromCore(data user.UserCore) User {
	return User{
		Name: data.Name,
		Email: data.Email,
		Password: data.Password,
		BookId: data.BookId,
		AuthorId: data.AuthorId,
	}
}

func toCore(data User) user.UserCore {
	return user.UserCore{
		ID: int(data.ID),
		Name: data.Name,
		Email: data.Email,
		Password: data.Password,
		BookId: data.BookId,
		AuthorId: data.AuthorId,
	}
}
