package delivery

import "github.com/iqbalpradipta/HTTP-RESTful-API-using-Go-and-a-SQL-Database/feature/user"

type userRequest struct {
	Name		string	`json:"name" form:"name"`
	Email	 	string 	`json:"email" form:"email" validate:"required,email"`
	Password 	string	`json:"password" form:"password" validate:"required"`
}

func toCore(data userRequest) user.UserCore {
	return user.UserCore{
		Name: data.Name,
		Email: data.Email,
		Password: data.Password,
	}
}

func toCoreLogin(data userRequest) user.UserCore {
	return user.UserCore{
		Email: data.Email,
		Password: data.Password,
	}
}