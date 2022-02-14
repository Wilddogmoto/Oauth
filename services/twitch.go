package services

import (
	"context"
	"fmt"
	"github.com/Wilddogmoto/Oauth/db"
	"github.com/Wilddogmoto/Oauth/logger"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/twitch"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type NewTwitchCredentials struct {
	oauth oauth2.Config
	token oauth2.Token
}

func TwitchCredentials() (ntc *NewTwitchCredentials, err error) {

	//clientId:= "6v8quhanak7fgmpkz97st515pmigsf"
	//secretCode:= "frc8br9txpir20unf1x231lcdt9a7a"
	//redirectURL:= "http://localhost:8989/auth/login_twitch"

	var (
		log = logger.Logger
		s   db.Services
	)

	err = db.DataBase.First(&s, "id = ?", "1").Error
	if err != nil {
		log.Errorf("error on database: %v", err)
		return
	}

	conf := oauth2.Config{
		ClientID:     s.ClientId,
		ClientSecret: s.SecretCode,
		Endpoint:     twitch.Endpoint,
		RedirectURL:  s.RedirectUri,
		Scopes:       []string{"user:read:email"},
	}

	ntc = &NewTwitchCredentials{
		oauth: conf,
	}

	return
}

func (ntc *NewTwitchCredentials) GetRedirect() (uri, state string, err error) {

	var (
		log        = logger.Logger
		u          *url.URL
		parameters = url.Values{}
	)

	if u, err = url.Parse(twitch.Endpoint.AuthURL); err != nil {
		log.Errorf("error on parsing url: %v", err)
		return
	}

	state = strconv.Itoa(int(time.Now().Unix()))

	parameters.Add("client_id", ntc.oauth.ClientID)
	parameters.Add("scope", strings.Join(ntc.oauth.Scopes, " "))
	parameters.Add("redirect_uri", ntc.oauth.RedirectURL)
	parameters.Add("response_type", "code")
	parameters.Add("state", state)

	u.RawQuery = parameters.Encode()

	fmt.Println(u.String())
	return u.String(), state, nil
}

func (ntc *NewTwitchCredentials) ExchangeCode(code string) (token *oauth2.Token, err error) {

	if token, err = ntc.oauth.Exchange(context.Background(), code); err != nil {
		return
	}

	return
}
