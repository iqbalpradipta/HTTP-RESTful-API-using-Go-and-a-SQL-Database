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

func (a *authorUseCase) GetAuthorById(id int) (data author.AuthorCore, err error) {
	data, err = a.authorData.SelectAuthorById(id)
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
	data, err := a.authorData.PutAuthor(id, update)
	if err != nil {
		return 0, errors.New("author failed to update")
	}
	return data, err
}

func (a *authorUseCase) DeleteAuthor(id int, delete author.AuthorCore) (int, error) {
	data, err := a.authorData.DelAuthor(id, delete)
	if err != nil {
		return data, errors.New("author failed to delete")
	}
	return data, err
}