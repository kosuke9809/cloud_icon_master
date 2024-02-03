package controller

import (
	"github.com/kosuke9809/cloud_icon_master/internal/usecase"
	"github.com/labstack/echo/v4"
)

type IOauthUserController interface {
	OauthLogin(c echo.Context) error
	Callback(c echo.Context) error
}

type oauthUserController struct {
	ouu usecase.IOauthUsecase
}

func NewOauthUserController(ouu usecase.IOauthUsecase) IOauthUserController {
	return &oauthUserController{ouu}
}

func (ouc *oauthUserController) OauthLogin(c echo.Context) error {
	url, err := ouc.ouu.GetLoginUrl()
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.Redirect(302, url)
}

func (ouc *oauthUserController) Callback(c echo.Context) error {
	code := c.QueryParam("code")
	userInfo, err := ouc.ouu.AuthenticateUser(c.Request().Context(), code)
	if err != nil {
		return c.JSON(400, err.Error())
	}

	return c.JSON(200, userInfo)
}
