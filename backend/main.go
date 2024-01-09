package main

import (
	"github.com/kosuke9809/aws_icons_quiz_backend/controller"
	"github.com/kosuke9809/aws_icons_quiz_backend/database"
	"github.com/kosuke9809/aws_icons_quiz_backend/oauth"
	"github.com/kosuke9809/aws_icons_quiz_backend/repository"
	"github.com/kosuke9809/aws_icons_quiz_backend/router"
	"github.com/kosuke9809/aws_icons_quiz_backend/usecase"
)

func main() {
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
