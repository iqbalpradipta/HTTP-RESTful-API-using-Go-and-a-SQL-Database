package delivery

import (
	"net/http"
	"strconv"

	"github.com/iqbalpradipta/HTTP-RESTful-API-using-Go-and-a-SQL-Database/feature/author"
	"github.com/iqbalpradipta/HTTP-RESTful-API-using-Go-and-a-SQL-Database/middleware"
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
	return c.JSON(http.StatusOK, helper.SuccessResponseDataHelper("success get data", fromCoreList(result)))
}

func (d *AuthorDelivery)GetAuthorById(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("error param"))
	}
	result, err := d.AuthorUseCase.GetAuthorById(idConv)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("error get param"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponseDataHelper("success get data", fromCore(result)))
}

func (d *AuthorDelivery)CreateAuthor(c echo.Context) error {
	var dataRequest authorRequest
	errBind := c.Bind(&dataRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("failed to Bind data"))
	}
	row, err := d.AuthorUseCase.CreateAuthor(toCore(dataRequest))
	if err != nil && row != 1 {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("failed to create data"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("success Create data"))
}

func (d *AuthorDelivery)UpdateAuthor(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("error param"))
	}

	idToken := middleware.ExtractToken(c)
	if idToken != idConv {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponseHelper("unauthorized"))
	}

	var dataUpdate authorRequest
	errBind := c.Bind(&dataUpdate)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("error bind data"))
	}

	row, err := d.AuthorUseCase.UpdateAuthor(idConv, toCore(dataUpdate))
	if err != nil || row == 0 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed to update data"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("Update success"))
}

func (d *AuthorDelivery)DeleteAuthor(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("error param")) 
	}
	var dataRemove authorRequest
	row, err := d.AuthorUseCase.DeleteAuthor(idConv, toCore(dataRemove))
	if err != nil || row == 0{
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("error delete data"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("success delete data"))
}