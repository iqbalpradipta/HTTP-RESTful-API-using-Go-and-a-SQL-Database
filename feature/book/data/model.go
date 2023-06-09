package data

import (
	"github.com/iqbalpradipta/HTTP-RESTful-API-using-Go-and-a-SQL-Database/feature/author/data"
	"github.com/iqbalpradipta/HTTP-RESTful-API-using-Go-and-a-SQL-Database/feature/book"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title           string
	PublishedYear 	string
	ISBN            int
	Authors         []data.Author `gorm:"many2many:author_books;"`
}

func (data *Book) toCore() book.BookCore {
	return book.BookCore{
		ID:             data.ID,
		Title:          data.Title,
		PublishedYear: 	data.PublishedYear,
		ISBN:           data.ISBN,
	}
}

func toCoreList(data []Book) []book.BookCore {
	var dataCore []book.BookCore
	for key := range data {
		dataCore = append(dataCore, data[key].toCore())
	}
	return dataCore
}

func fromCore(dataCore book.BookCore) Book {
	return Book{
		Title:           dataCore.Title,
		PublishedYear: 	 dataCore.PublishedYear,
		ISBN:            dataCore.ISBN,
	}
}
