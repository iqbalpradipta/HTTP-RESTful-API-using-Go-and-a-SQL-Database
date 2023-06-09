package delivery

import "github.com/iqbalpradipta/HTTP-RESTful-API-using-Go-and-a-SQL-Database/feature/book"

type bookRequest struct {
	Title			string	`json:"title" form:"title"`
	PublishedYear	string	`json:"publishedyear" form:"publishedyear"`
	ISBN			int		`json:"isbn" form:"isbn"`
}

func toCore(data bookRequest) book.BookCore {
	return book.BookCore{
		Title: data.Title,
		PublishedYear: data.PublishedYear,
		ISBN: data.ISBN,
	}
}