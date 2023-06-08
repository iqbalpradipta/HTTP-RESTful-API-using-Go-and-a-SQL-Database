package delivery

import "github.com/iqbalpradipta/HTTP-RESTful-API-using-Go-and-a-SQL-Database/feature/author"

type authorRequest struct {
	Name		string
	Country		string
}

func toCore(data authorRequest) author.AuthorCore {
	return author.AuthorCore{
		Name: data.Name,
		Country: data.Country,
	}
}