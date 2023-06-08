package data

import (
	"errors"

	"github.com/iqbalpradipta/HTTP-RESTful-API-using-Go-and-a-SQL-Database/feature/author"
	"gorm.io/gorm"
)

type authorData struct {
	db *gorm.DB
}

func NewQuery(db *gorm.DB) author.IdataInterface {
	return &authorData{
		db: db,
	}
}

func (r *authorData)SelectAllAuthor() ([]author.AuthorCore, error) {
	var data []Author
	tx := r.db.Find(&data)
	if tx.Error != nil {
		return nil, tx.Error
	}

	dataCore := toCoreList(data)
	return dataCore, nil
}

func (r *authorData)SelectAuthorByName(name string) (author.AuthorCore, error)  {
	var data Author
	tx := r.db.First(&data, name)
	if tx.Error != nil {
		return author.AuthorCore{}, tx.Error
	}
	return data.toCore(), nil
}

func (r *authorData)InsertAuthor(data author.AuthorCore) (int, error) {
	dataModel := fromCore(data)
	tx := r.db.Create(&dataModel)
	if tx.Error != nil {
		return 0, nil
	}
	return int(tx.RowsAffected), nil
}

func (r *authorData)PutAuthor(id int , update author.AuthorCore) (int, error)  {
	tx := r.db.Model(&Author{}).Where("id = ?", id).Updates(fromCore(update))
	if tx.Error != nil {
		return -1, tx.Error
	}
	
	if tx.RowsAffected == 0 {
		return 0, errors.New("no data update")
	}

	return int(tx.RowsAffected), nil
}

func (r *authorData)DelAuthor(id int, delete author.AuthorCore) (int, error)  {
	tx := r.db.Where("id = ?", id).Delete(&Author{})
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("failed to delete data")
	}
	return int(tx.RowsAffected), nil
}