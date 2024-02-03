package api

import (
	"fmt"

	"github.com/kosuke9809/cloud_icon_master/internal/domain/repository"
	database "github.com/kosuke9809/cloud_icon_master/internal/infra/database/config"
	"github.com/kosuke9809/cloud_icon_master/internal/infra/oauth"
	"github.com/kosuke9809/cloud_icon_master/internal/interfaces/controller"
	"github.com/kosuke9809/cloud_icon_master/internal/interfaces/router"
	"github.com/kosuke9809/cloud_icon_master/internal/usecase"
)

func Start() {
	fmt.Println("Start Backend Server")
	db := database.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)

	iconRepository := repository.NewIconRepository(db)
	iconUsecase := usecase.NewIconUsecase(iconRepository)
	iconController := controller.NewIconController(iconUsecase)

	oauthConfig := oauth.GetConnect()
	oauthUserRepository := repository.NewOauthUserRepository(db, oauthConfig)
	oauthUserUsecase := usecase.NewOauthUsecase(oauthUserRepository)
	oauthUserController := controller.NewOauthUserController(oauthUserUsecase)

	e := router.NewRouter(userController, iconController, oauthUserController)
	e.Logger.Fatal(e.Start(":8080"))
}
