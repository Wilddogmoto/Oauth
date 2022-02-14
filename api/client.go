package api

import (
	"github.com/Wilddogmoto/Oauth/configuration"
	"github.com/Wilddogmoto/Oauth/logger"
	"github.com/gin-gonic/gin"
)

func RouterClient() {

	var (
		log = logger.Logger
	)

	log.Infof("initialisation router client on port: %v >>>", configuration.Parameters.Common.ApiPort)

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	privat := router.Group("/auth")
	{
		privat.POST("/twitch", OauthLoginTwitch)
		privat.GET("/login_twitch", GetOauthTwitch)
	}

	if err := router.Run(configuration.Parameters.Common.ApiPort); err != nil {
		log.Fatalf("route start error on RouterClient: %v", err)
		return
	}
}
