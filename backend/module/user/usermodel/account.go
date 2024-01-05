package usermodel

import (
	"backend/component/tokenprovider"
)

type Account struct {
	AccessToken  *tokenprovider.Token `json:"accessToken"`
	RefreshToken *tokenprovider.Token `json:"refreshToken"`
}

func NewAccount(at, rt *tokenprovider.Token) *Account {
	return &Account{
		AccessToken:  at,
		RefreshToken: rt,
	}
}
