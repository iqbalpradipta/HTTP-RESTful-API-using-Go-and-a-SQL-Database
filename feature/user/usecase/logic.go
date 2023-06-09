package usecase

import (
	"github.com/iqbalpradipta/HTTP-RESTful-API-using-Go-and-a-SQL-Database/feature/user"
	"github.com/iqbalpradipta/HTTP-RESTful-API-using-Go-and-a-SQL-Database/middleware"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userData	user.IdataInterface
}

func NewLogic(data user.IdataInterface) user.IusecaseInterface  {
	return &userUsecase{
		userData: data,
	}
}

func (u *userUsecase) CreateUser(data user.UserCore)(int, error) {
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, nil
	}
	data.Password = string(passwordHashed)

	row, err := u.userData.RegisterUser(data)
	return row, err
}

func (u *userUsecase) LoginUser(data user.UserCore) (token string, err error)  {
	result, err := u.userData.AuthUser(data.Email)

	if err != nil {
		return token, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(data.Password))

	if err != nil {
		return token, err
	}

	return middleware.CreateToken(result.ID)
}

