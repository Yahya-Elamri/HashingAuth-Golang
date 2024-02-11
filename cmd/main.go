package main

import (
	"fmt"
	"net/http"

	"github.com/Yahya-Elamri/hashingauth/handler"
	"github.com/Yahya-Elamri/hashingauth/utils"
	"github.com/Yahya-Elamri/hashingauth/view/form"
	"github.com/Yahya-Elamri/hashingauth/view/home"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	utils.Db = utils.ConnectDb()

	e.POST("/addnewuser", handler.AddUser)
	e.POST("/checkuser", handler.CheckUsername)
	e.POST("/checkemail", handler.CheckEmail)
	e.POST("/userauth", handler.UserAuth)

	e.GET("/allusers", func(c echo.Context) error {
		var users []utils.UserData
		utils.Db.Raw("SELECT * FROM auth.users;").Scan(&users)
		return c.JSON(200, users)
	})
	e.GET("/", func(c echo.Context) error {
		var user utils.UserData
		session, err := session.Get("user", c)
		if err != nil {
			panic(err)
		}
		Id := session.Values["user_id"]
		if Id == nil {
			return c.Redirect(307, "/login")
		}
		utils.Db.Raw(fmt.Sprintf("SELECT * FROM `auth`.`users` WHERE idusers='%d'", Id)).Scan(&user)
		return home.Show(user).Render(c.Request().Context(), c.Response())
	})
	e.GET("/disconnect", func(c echo.Context) error {
		session, _ := session.Get("user", c)

		// Clear session data
		delete(session.Values, "user_id")
		session.Save(c.Request(), c.Response())
		c.Response().Header().Set("HX-Redirect", "/")
		return c.String(http.StatusOK, "disconnecting successful! Redirecting...")
	})
	e.GET("/login", func(c echo.Context) error {
		return form.LogInForm().Render(c.Request().Context(), c.Response())
	})
	e.GET("/signup", func(c echo.Context) error {
		return form.SignUpForm().Render(c.Request().Context(), c.Response())
	})
	e.Logger.Fatal(e.Start(":3000"))
}
