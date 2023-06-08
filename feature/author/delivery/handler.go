package delivery

import (
	"net/http"

	"github.com/iqbalpradipta/HTTP-RESTful-API-using-Go-and-a-SQL-Database/feature/author"
	"github.com/iqbalpradipta/HTTP-RESTful-API-using-Go-and-a-SQL-Database/utils/helper"
	"github.com/labstack/echo/v4"
)

type AuthorDelivery struct {
	AuthorUseCase	author.IusecaseInterface
}

func NewHandler(e *echo.Echo, useCase author.IusecaseInterface){
	handler := &AuthorDelivery{
		AuthorUseCase: useCase,
	}

	e.GET("/author", handler.GetAllAuthor)

}

func (d *AuthorDelivery) GetAllAuthor(c echo.Context) error {
	result, err := d.AuthorUseCase.GetAllAuthor()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed get data"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponseDataHelper("success get data", result))
}