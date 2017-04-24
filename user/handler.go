package user

import (
	"net/http"

	"github.com/labstack/echo"
)

//CreateUser : create user
func CreateUser(c echo.Context) error {
	user := new(User)
	//var err error
	return c.JSON(http.StatusCreated, user)
}
