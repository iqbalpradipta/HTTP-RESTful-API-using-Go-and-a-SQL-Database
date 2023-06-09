package factory

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	authorData "github.com/iqbalpradipta/HTTP-RESTful-API-using-Go-and-a-SQL-Database/feature/author/data"
	authorDelivery "github.com/iqbalpradipta/HTTP-RESTful-API-using-Go-and-a-SQL-Database/feature/author/delivery"
	authorUsecase "github.com/iqbalpradipta/HTTP-RESTful-API-using-Go-and-a-SQL-Database/feature/author/usecase"

	bookData "github.com/iqbalpradipta/HTTP-RESTful-API-using-Go-and-a-SQL-Database/feature/book/data"
	bookDelivery "github.com/iqbalpradipta/HTTP-RESTful-API-using-Go-and-a-SQL-Database/feature/book/delivery"
	bookUsecase "github.com/iqbalpradipta/HTTP-RESTful-API-using-Go-and-a-SQL-Database/feature/book/usecase"

	userData "github.com/iqbalpradipta/HTTP-RESTful-API-using-Go-and-a-SQL-Database/feature/user/data"
	userDelivery "github.com/iqbalpradipta/HTTP-RESTful-API-using-Go-and-a-SQL-Database/feature/user/delivery"
	userUsecase "github.com/iqbalpradipta/HTTP-RESTful-API-using-Go-and-a-SQL-Database/feature/user/usecase"
)

func InitFactory(e *echo.Echo, db *gorm.DB)  {
	authorDataFactory := authorData.NewQuery(db)
	authorUsecaseFactory := authorUsecase.NewLogic(authorDataFactory)
	authorDelivery.NewHandler(e, authorUsecaseFactory)

	bookDataFactory := bookData.NewQuery(db)
	bookUsecaseFactory := bookUsecase.NewLogic(bookDataFactory)
	bookDelivery.NewHandler(e, bookUsecaseFactory)

	userDataFactory := userData.NewQuery(db)
	userUsecaseFactory := userUsecase.NewLogic(userDataFactory)
	userDelivery.NewHandler(e, userUsecaseFactory)
}