package controller

import (
	"net/http"

	"github.com/kosuke9809/cloud_icon_master/internal/usecase"
	"github.com/labstack/echo/v4"
)

type IIconController interface {
	GetAllIcons(c echo.Context) error
	GetRandomIcons(c echo.Context) error
}

type iconController struct {
	iu usecase.IIconUsecase
}

func NewIconController(iu usecase.IIconUsecase) IIconController {
	return &iconController{iu}
}

func (ic *iconController) GetAllIcons(c echo.Context) error {
	iconRes, err := ic.iu.GetAllIcons()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, iconRes)
}

func (ic *iconController) GetRandomIcons(c echo.Context) error {
	iconRes, err := ic.iu.GetRandomIcons(4)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, iconRes)
}
