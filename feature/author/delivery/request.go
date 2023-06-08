package delivery

import "github.com/iqbalpradipta/HTTP-RESTful-API-using-Go-and-a-SQL-Database/feature/author"

type AuthorRequest struct {
	Name		string `json:"name" form:"name"`
	Country		string `json:"country" form:"country"`
}

func toCore(data AuthorRequest) author.AuthorCore {
	return author.AuthorCore{
		Name: data.Name,
		Country: data.Country,
	}
}