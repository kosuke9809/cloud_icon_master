package router

import (
	"net/http"
	"os"

	"github.com/kosuke9809/aws_icons_quiz_backend/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(uc controller.IUserController, ic controller.IIconController, ouc controller.IOauthUserController) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{os.Getenv("FE_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	e.GET("/icons", ic.GetAllIcons)
	e.GET("/quiz", ic.GetRandomIcons)
	e.GET("/oauth/login", ouc.OauthLogin)
	e.GET("/callback", ouc.Callback)
	return e
}
