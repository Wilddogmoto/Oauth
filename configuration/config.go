package configuration

import (
	"encoding/json"
	"github.com/Wilddogmoto/Oauth/logger"
	"os"
)

type (
	Config struct {
		DB struct {
			Server   string `json:"server"`
			User     string `json:"user"`
			Password string `json:"password"`
			Dbname   string `json:"dbname"`
		}`json:"db"`

		Common struct {
			ApiPort string `json:"api_port"`
		} `json:"common"`
	}
)

var Parameters *Config


func PreloadConfig() {

	var (
		log  = logger.Logger
		err  error
		file *os.File
	)

	if file, err = os.Open("./config.json"); err != nil {
		log.Fatalf("err on open config file:%v", err)
	}

	if err = json.NewDecoder(file).Decode(&Parameters); err != nil {
		log.Fatalf("error decoding config file:%v", err)
	}
}
