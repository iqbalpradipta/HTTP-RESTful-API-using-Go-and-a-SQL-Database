package usecase

import (
	"errors"

	"github.com/iqbalpradipta/HTTP-RESTful-API-using-Go-and-a-SQL-Database/feature/author"
)

type authorUseCase struct{
	authorData author.IdataInterface
}

func NewLogic(data author.IdataInterface) author.IusecaseInterface {
	return &authorUseCase {
		authorData: data,
	}
}

func (a *authorUseCase) GetAllAuthor() ([]author.AuthorCore, error) {
	result, err := a.authorData.SelectAllAuthor()
	return result, err
}

func (a *authorUseCase) GetAuthorByName(name string) (data author.AuthorCore, err error) {
	data, err = a.authorData.SelectAuthorByName(name)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (a *authorUseCase) CreateAuthor(data author.AuthorCore) (int, error) {
	row, err := a.authorData.InsertAuthor(data)
	if err != nil {
		return row, err
	}
	return row, nil
}

func (a *authorUseCase) UpdateAuthor(id int, update author.AuthorCore) (int, error) {
	if update.ID == 0 {
		return -1, errors.New("id not found")
	}

	data, err := a.authorData.PutAuthor(id, update)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (a *authorUseCase) DeleteAuthor(id int, delete author.AuthorCore) (int, error) {
	if delete.ID == 0 {
		return -1, errors.New("id not found")
	}
	
	data, err := a.authorData.DelAuthor(id, delete)
	if err != nil {
		return data, err
	}
	return data, nil
}