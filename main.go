package main

import (
	"github.com/Wilddogmoto/Oauth/api"
	"github.com/Wilddogmoto/Oauth/configuration"
	"github.com/Wilddogmoto/Oauth/db"
	"github.com/Wilddogmoto/Oauth/logger"
)

func main() {

	configuration.PreloadConfig() //load config

	logger.InitLogger() // init logger

	db.DataBaseConnected() // init database

	api.RouterClient() // init router
}
