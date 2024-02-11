package handler

import (
	"fmt"

	"github.com/Yahya-Elamri/hashingauth/utils"
	"github.com/labstack/echo/v4"
)

func AddUser(c echo.Context) error {
	var user utils.UserData
	Username := c.FormValue("username")
	Email := c.FormValue("email")
	Password := c.FormValue("password")
	hashedpassword, _ := utils.Hashing(Password)
	str := fmt.Sprintf("INSERT INTO `auth`.`users` (`username`, `email`, `password`) VALUES ('%s', '%s', '%s');", Username, Email, hashedpassword)
	err := utils.Db.Raw(str).Scan(&user).Error
	if err != nil {
		return c.HTML(404, "<p>something went wrong</p>")
	}
	return c.HTML(200, "<div> you have been successfully signed up <br> <center><a href='/login'> Go to Login page </a></center> <div/>")
}
