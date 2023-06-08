package delivery

import "github.com/iqbalpradipta/HTTP-RESTful-API-using-Go-and-a-SQL-Database/feature/author"

type authorResponse struct {
	Name		string	`json:"name"`
	Country		string	`json:"country"`
}

func fromCore(data author.AuthorCore) authorResponse {
	return authorResponse{
		Name: data.Name,
		Country: data.Country,
	}
}

func fromCoreList(data []author.AuthorCore) []authorResponse {
	var dataResponse []authorResponse
	for _, v := range data {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}