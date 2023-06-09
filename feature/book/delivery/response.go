package delivery

import "github.com/iqbalpradipta/HTTP-RESTful-API-using-Go-and-a-SQL-Database/feature/book"

type bookResponse struct {
	ID					uint	`json:"id"`
	Title				string	`json:"title"`
	PublishedYear		string	`json:"publishedyear"`
	ISBN				int		`json:"isbn"`
}

func fromCore(data book.BookCore) bookResponse {
	return bookResponse{
		ID: data.ID,
		Title: data.Title,
		PublishedYear: data.PublishedYear,
		ISBN: data.ISBN,
	}
}

func fromCoreList(data []book.BookCore) []bookResponse {
	var dataResponse []bookResponse
	for _, v := range data {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}