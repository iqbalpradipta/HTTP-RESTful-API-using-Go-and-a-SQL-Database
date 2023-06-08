package factory

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	authorData "github.com/iqbalpradipta/HTTP-RESTful-API-using-Go-and-a-SQL-Database/feature/author/data"
	authorDelivery "github.com/iqbalpradipta/HTTP-RESTful-API-using-Go-and-a-SQL-Database/feature/author/delivery"
	authorUsecase "github.com/iqbalpradipta/HTTP-RESTful-API-using-Go-and-a-SQL-Database/feature/author/usecase"
)

func InitFactory(e *echo.Echo, db *gorm.DB)  {
	authorDataFactory := authorData.NewQuery(db)
	authorUsecaseFactory := authorUsecase.NewLogic(authorDataFactory)
	authorDelivery.NewHandler(e, authorUsecaseFactory)
}