package delivery

import (
	"net/http"
	"strconv"

	"github.com/iqbalpradipta/HTTP-RESTful-API-using-Go-and-a-SQL-Database/feature/book"
	"github.com/iqbalpradipta/HTTP-RESTful-API-using-Go-and-a-SQL-Database/middleware"
	"github.com/iqbalpradipta/HTTP-RESTful-API-using-Go-and-a-SQL-Database/utils/helper"
	"github.com/labstack/echo/v4"
)

type BookDelivery struct {
	BookUseCase	book.IusecaseInterface
}

func NewHandler(e *echo.Echo, useCase book.IusecaseInterface){
	handler := &BookDelivery{
		BookUseCase: useCase,
	}

	e.GET("/book", handler.GetAllBook, middleware.JWTMiddleware())
	e.GET("/book/:id", handler.GetBookById, middleware.JWTMiddleware())
	e.POST("/book", handler.CreateBook, middleware.JWTMiddleware())
	e.PUT("/book/:id", handler.UpdateBook, middleware.JWTMiddleware())
	e.DELETE("/book/:id", handler.DeleteBook, middleware.JWTMiddleware())

}

func (d *BookDelivery) GetAllBook(c echo.Context) error {
	result, err := d.BookUseCase.GetAllBook()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed get data"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponseDataHelper("success get data", fromCoreList(result)))
}

func (d *BookDelivery)GetBookById(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("error param"))
	}
	result, err := d.BookUseCase.GetBookById(idConv)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("error get data"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponseDataHelper("success get data", fromCore(result)))
}

func (d *BookDelivery)CreateBook(c echo.Context) error {
	var dataRequest bookRequest
	errBind := c.Bind(&dataRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("failed to Bind data"))
	}
	row, err := d.BookUseCase.CreateBook(toCore(dataRequest))
	if err != nil && row != 1 {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("failed to create data"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("success Create data"))
}

func (d *BookDelivery)UpdateBook(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil || idConv == 0{
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("error update by param"))
	}

	idToken := middleware.ExtractToken(c)
	if idToken != idConv {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponseHelper("unauthorized"))
	}

	var dataUpdate bookRequest
	errBind := c.Bind(&dataUpdate)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("error bind data"))
	}

	row, err := d.BookUseCase.UpdateBook(idConv, toCore(dataUpdate))
	if err != nil || row == 0 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed to update data"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("Update success"))
}

func (d *BookDelivery)DeleteBook(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("error param")) 
	}
	var dataRemove bookRequest
	row, err := d.BookUseCase.DeleteBook(idConv, toCore(dataRemove))
	if err != nil || row == 0{
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("error delete data"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("success delete data"))
}