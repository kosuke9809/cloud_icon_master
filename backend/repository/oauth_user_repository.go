package repository

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/kosuke9809/aws_icons_quiz_backend/model"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

type IOauthUserRepository interface {
	CreateOauthUser(ctx context.Context, user *model.GoogleUser) error
	GetUserInfo(ctx context.Context, token *oauth2.Token) (*model.GoogleUser, error)
	ExchangeToken(ctx context.Context, token string) (*oauth2.Token, error)
	GetLoginUrl() (string, error)
}

type oauthUserRepository struct {
	db     *gorm.DB
	config *oauth2.Config
}

func NewOauthUserRepository(db *gorm.DB, config *oauth2.Config) IOauthUserRepository {
	return &oauthUserRepository{db, config}
}

func (our *oauthUserRepository) CreateOauthUser(ctx context.Context, user *model.GoogleUser) error {
	if err := our.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// Google OAuth2.0を使用してユーザー情報を取得し、Emailを元にDBからユーザー情報を取得する関数
func (our *oauthUserRepository) GetUserInfo(ctx context.Context, token *oauth2.Token) (*model.GoogleUser, error) {
	client := our.config.Client(ctx, token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	googleUser := model.GoogleUser{}
	if err := json.NewDecoder(resp.Body).Decode(&googleUser); err != nil {
		return nil, err
	}

	if err := our.db.Where("email=?", googleUser.Email).First(&googleUser).Error; err != nil {
		// ユーザーが存在しない場合は新規作成
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := our.CreateOauthUser(ctx, &googleUser); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	return &googleUser, nil
}

func (our *oauthUserRepository) ExchangeToken(ctx context.Context, code string) (*oauth2.Token, error) {
	return our.config.Exchange(ctx, code)
}

func (our *oauthUserRepository) GetLoginUrl() (string, error) {
	state := "state"
	url := our.config.AuthCodeURL(state, oauth2.AccessTypeOffline)
	return url, nil
}
