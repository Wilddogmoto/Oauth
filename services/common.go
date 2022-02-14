package services

import "golang.org/x/oauth2"

type Service interface {
	GetRedirect() (string, string, error)
	ExchangeCode(string) (*oauth2.Token, error)
}
