package usecase

import (
	"context"

	"github.com/kosuke9809/aws_icons_quiz_backend/model"
	"github.com/kosuke9809/aws_icons_quiz_backend/repository"
)

type oauthUserUsecase struct {
	our repository.IOauthUserRepository
}

type IOauthUsecase interface {
	AuthenticateUser(ctx context.Context, code string) (*model.GoogleUser, error)
	GetLoginUrl() (string, error)
}

func NewOauthUsecase(our repository.IOauthUserRepository) IOauthUsecase {
	return &oauthUserUsecase{our}
}

func (ouu *oauthUserUsecase) AuthenticateUser(ctx context.Context, code string) (*model.GoogleUser, error) {
	token, err := ouu.our.ExchangeToken(ctx, code)
	if err != nil {
		return nil, err
	}

	userInfo, err := ouu.our.GetUserInfo(ctx, token)
	if err != nil {
		return nil, err
	}

	return userInfo, nil
}

func (ouu *oauthUserUsecase) GetLoginUrl() (string, error) {
	return ouu.our.GetLoginUrl()
}
