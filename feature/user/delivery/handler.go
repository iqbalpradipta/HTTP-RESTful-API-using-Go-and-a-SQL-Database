package delivery

import (
	"net/http"

	"github.com/iqbalpradipta/HTTP-RESTful-API-using-Go-and-a-SQL-Database/feature/user"
	"github.com/iqbalpradipta/HTTP-RESTful-API-using-Go-and-a-SQL-Database/utils/helper"
	"github.com/labstack/echo/v4"
)

type UserDelivery struct {
	userUseCase user.IusecaseInterface
}

func NewHandler(e *echo.Echo, useCase user.IusecaseInterface) {
	handler := &UserDelivery{
		userUseCase: useCase,
	}

	e.POST("/register", handler.CreateUser)
	e.POST("/login", handler.LoginUser)
}


func (d *UserDelivery)CreateUser(c echo.Context) error {
	var dataRequest userRequest
	errBind := c.Bind(&dataRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("failed to Bind data"))
	}
	row, err := d.userUseCase.CreateUser(toCore(dataRequest))
	if err != nil && row != 1 {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("failed to create data"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("success Create data"))
}

func (d *UserDelivery)LoginUser(c echo.Context) error {
	loginRequest := userRequest{}
	c.Bind(&loginRequest)

	err := c.Validate(loginRequest)
	if err != nil {
		return c.JSON(http.StatusBadGateway, helper.FailedResponseHelper(err.Error()))
	}

	loginCore := toCoreLogin(loginRequest)
	token, err := d.userUseCase.LoginUser(loginCore)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponseDataHelper("success logged", map[string]interface{}{
		"token": token,
	}))

}
