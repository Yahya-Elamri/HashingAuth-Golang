package handler

import (
	"fmt"
	"net/http"

	"github.com/Yahya-Elamri/hashingauth/utils"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func UserAuth(c echo.Context) error {
	var User utils.UserData
	Username := c.FormValue("username")
	Password := c.FormValue("password")
	SqlQuery := fmt.Sprintf("SELECT * FROM `auth`.`users` WHERE username='%s' OR email='%s'", Username, Username)
	utils.Db.Raw(SqlQuery).Scan(&User)
	if User.Username == "" {
		return c.HTML(401, "<p class='text-sm text-bold text-red-500'> Username or Email is incorrect</p>")
	}
	if !utils.CompairHash(Password, User.Password) {
		return c.HTML(402, "<p class='text-sm text-bold text-red-500'>password incorrect</p>")
	}
	session, _ := session.Get("user", c)
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	session.Values["user_id"] = User.Idusers
	session.Save(c.Request(), c.Response())
	c.Response().Header().Set("HX-Redirect", "/")
	return c.String(http.StatusOK, "Login successful! Redirecting...")
}
