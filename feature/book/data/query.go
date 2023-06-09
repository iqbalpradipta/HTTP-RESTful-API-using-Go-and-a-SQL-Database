package data

import (
	"errors"

	"github.com/iqbalpradipta/HTTP-RESTful-API-using-Go-and-a-SQL-Database/feature/book"
	"gorm.io/gorm"
)

type bookData struct {
	db *gorm.DB
}

func NewQuery(db *gorm.DB) book.IdataInterface {
	return &bookData{
		db: db,
	}
}

func (r *bookData)SelectAllBook() ([]book.BookCore, error) {
	var data []Book
	tx := r.db.Find(&data)
	if tx.Error != nil {
		return nil, tx.Error
	}

	dataCore := toCoreList(data)
	return dataCore, nil
}

func (r *bookData)SelectBookById(id int) (book.BookCore, error)  {
	var data Book
	tx := r.db.First(&data, id)
	if tx.Error != nil {
		return book.BookCore{}, tx.Error
	}
	return data.toCore(), nil
}

func (r *bookData)InsertBook(data book.BookCore) (int, error) {
	dataModel := fromCore(data)
	tx := r.db.Create(&dataModel)
	if tx.Error != nil {
		return 0, nil
	}
	return int(tx.RowsAffected), nil
}

func (r *bookData)PutBook(id int , update book.BookCore) (int, error)  {
	tx := r.db.Model(&Book{}).Where("id = ?", id).Updates(fromCore(update))
	if tx.Error != nil {
		return -1, tx.Error
	}
	
	if tx.RowsAffected == 0 {
		return 0, errors.New("no data update")
	}

	return int(tx.RowsAffected), nil
}

func (r *bookData)DelBook(id int, delete book.BookCore) (int, error)  {
	tx := r.db.Where("id = ?", id).Delete(&Book{})
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("failed to delete data")
	}
	return int(tx.RowsAffected), nil
}