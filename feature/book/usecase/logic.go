package usecase

import (
	"errors"

	"github.com/iqbalpradipta/HTTP-RESTful-API-using-Go-and-a-SQL-Database/feature/book"
)

type bookUseCase struct{
	bookData book.IdataInterface
}

func NewLogic(data book.IdataInterface) book.IusecaseInterface {
	return &bookUseCase {
		bookData: data,
	}
}

func (a *bookUseCase) GetAllBook() ([]book.BookCore, error) {
	result, err := a.bookData.SelectAllBook()
	return result, err
}

func (a *bookUseCase) GetBookById(id int) (data book.BookCore, err error) {
	data, err = a.bookData.SelectBookById(id)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (a *bookUseCase) CreateBook(data book.BookCore) (int, error) {
	row, err := a.bookData.InsertBook(data)
	if err != nil {
		return row, err
	}
	return row, nil
}

func (a *bookUseCase) UpdateBook(id int, update book.BookCore) (int, error) {
	data, err := a.bookData.PutBook(id, update)
	if err != nil {
		return 0, errors.New("book failed to update")
	}
	return data, err
}

func (a *bookUseCase) DeleteBook(id int, delete book.BookCore) (int, error) {
	data, err := a.bookData.DelBook(id, delete)
	if err != nil {
		return data, errors.New("book failed to delete")
	}
	return data, err
}