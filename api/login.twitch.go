package api

import (
	"errors"
	"github.com/Wilddogmoto/Oauth/db"
	"github.com/Wilddogmoto/Oauth/logger"
	"github.com/Wilddogmoto/Oauth/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type loginOauth struct {
	Code  string `json:"code"`
	State string `json:"state"`
}

func OauthLoginTwitch(ctx *gin.Context) {

	var (
		log   = logger.Logger
		uri   string
		state string
		err   error
		ntc   services.Service
	)

	if ntc, err = services.TwitchCredentials(); err != nil {
		log.Errorf("error on database: %v", err)
		sendResponse(10, nil, ctx)
		return
	}

	if uri, state, err = ntc.GetRedirect(); err != nil {
		log.Errorf("error on GetRedirect: %v", err)
		sendResponse(10, nil, ctx)
		return
	}

	entry := db.ServiceLogins{
		State: state,
	}

	if err = db.DataBase.Create(&entry).Error; err != nil {
		log.Errorf("error created on data base: %v", err)
		sendResponse(10, nil, ctx)
		return
	}

	data := make(map[string]string)
	data["redirect_url"] = uri
	data["state"] = state
	sendResponse(0, data, ctx)
	return
}

func GetOauthTwitch(ctx *gin.Context) {

	var (
		log = logger.Logger
		sl  db.ServiceLogins
		err error

		ntc services.Service
	)

	if ntc, err = services.TwitchCredentials(); err != nil {
		log.Errorf("error on database: %v", err)
		sendResponse(10, nil, ctx)
		return
	}

	err = db.DataBase.First(&sl, "state = ?", ctx.Query("state")).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Errorf("error on data base: %v", err)
		sendResponse(10, nil, ctx)
		return
	}

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		log.Errorf("state not found: %v", err)
		sendResponse(10, nil, ctx)
		return
	}

	if _, err = ntc.ExchangeCode(ctx.Query("code")); err != nil {
		log.Errorf("error on getting token from social network: %v", err)
		sendResponse(2, nil, ctx)
		return
	}
	ctx.Request.URL.Query().Get("scope")
	ctx.Param("email")

	//todo
	sendResponse(0, nil, ctx)

}
