package handler

import (
	"awesomeProject/model"
	req2 "awesomeProject/model/req"
	echo2 "github.com/labstack/echo"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
}

func (u *UserHandler) HandleSignUp(c echo2.Context) error {
	req := req2.ReqSignUp{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	type User struct {
		Email    string
		FullName string
	}
	user := User{
		Email:    "nda2397@gmail.com",
		FullName: "Nguyễn Duy Anh",
	}
	return c.JSON(http.StatusOK, user)
}

func (u *UserHandler) HandleSignIn(c echo2.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"user":  "DuyAnh",
		"email": "nda2397@gmail.com",
	})
}
