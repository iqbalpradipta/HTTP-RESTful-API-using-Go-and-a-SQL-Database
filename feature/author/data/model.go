package data

import (
	"github.com/iqbalpradipta/HTTP-RESTful-API-using-Go-and-a-SQL-Database/feature/author"
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name		string
	Country		string
}

func (data *Author)toCore() author.AuthorCore  {
	return author.AuthorCore{
		ID: data.ID,
		Name: data.Name,
		Country: data.Country,
	}
}

func toCoreList(data []Author) []author.AuthorCore {
	var dataCore []author.AuthorCore
	for key := range data {
		dataCore = append(dataCore, data[key].toCore())
	}
	return dataCore
}

func fromCore(dataCore author.AuthorCore) Author {
	return Author{
		Name: dataCore.Name,
		Country: dataCore.Country,
	}
}


