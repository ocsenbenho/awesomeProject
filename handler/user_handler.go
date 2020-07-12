package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)
func HandleSignIn(c echo.Context) error{
	return c.JSON(http.StatusOK, echo.Map{
		"user":"DuyAnh",
		"email":"nda2397@gmail.com",
	})
}
func HandleSignUp( c echo.Context)error{
	type User struct{
		Email string
		FullName string
	}
	user := User{
		Email:"nda2397@gmail.com",
		FullName: "Nguyễn Duy Anh",
	}
	return c.JSON(http.StatusOK, user)
}
