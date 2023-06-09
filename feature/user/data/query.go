package data

import (
	"github.com/iqbalpradipta/HTTP-RESTful-API-using-Go-and-a-SQL-Database/feature/user"
	"gorm.io/gorm"
)

type userData struct {
	db *gorm.DB
}

func NewQuery(db *gorm.DB) user.IdataInterface {
	return &userData{
		db: db,
	}
}

func (r *userData)RegisterUser(data user.UserCore) (int, error)  {
	dataModel := fromCore(data)
	tx := r.db.Create(&dataModel)
	if tx.Error != nil {
		return 0, nil
	}
	return int(tx.RowsAffected), nil
}

func (r *userData) AuthUser(email string) (user.UserCore, error) {
	data := User{}
	tx := r.db.Model(&User{}).Where("email = ?", email).First(&data)
	if tx.Error != nil {
		return user.UserCore{}, tx.Error
	}

	var dataUser = toCore(data)
	return dataUser, nil
}