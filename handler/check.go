package handler

import (
	"fmt"

	"github.com/Yahya-Elamri/hashingauth/utils"
	"github.com/labstack/echo/v4"
)

func CheckUsername(c echo.Context) error {
	var User utils.UserData
	Username := c.FormValue("username")
	username := fmt.Sprintf("SELECT * FROM `auth`.`users` WHERE username='%s';", Username)
	utils.Db.Raw(username).Scan(&User)
	if User.Username != "" {
		return c.HTML(401, fmt.Sprintf("<p class='text-sm text-bold text-red-500'>%s already exists</p>", Username))
	}
	return c.HTML(301, "<p class='text-sm text-bold text-green-500'>Good</p>")
}

func CheckEmail(c echo.Context) error {
	var User utils.UserData
	Email := c.FormValue("email")
	email := fmt.Sprintf("SELECT * FROM `auth`.`users` WHERE email='%s';", Email)
	utils.Db.Raw(email).Scan(&User)
	if User.Email != "" {
		return c.HTML(402, "<p class='text-sm text-bold text-red-500'>Email already exists</p>")
	}
	return c.HTML(302, "<p class='text-sm text-bold text-green-500'>Good</p>")
}
