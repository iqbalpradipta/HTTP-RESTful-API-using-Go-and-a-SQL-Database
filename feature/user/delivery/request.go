package delivery

import "github.com/iqbalpradipta/HTTP-RESTful-API-using-Go-and-a-SQL-Database/feature/user"

type userRequest struct {
	Email	 	string 	`json:"email" form:"email" validate:"required,email"`
	Password 	string	`json:"password" form:"password" validate:"require"`
}

func toCore(data userRequest) user.UserCore {
	return user.UserCore{
		Email: data.Email,
		Password: data.Password,
	}
}