package oauth2

import (
	"encoding/json"
	"github.com/1851616111/util/http"
)

//docs https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140842&token=&lang=zh_CN
type Config struct {
	AppID  string
	Secret string
}

func NewTokenConfig(app_id, secret string) *Config {
	return &Config{
		AppID:  app_id,
		Secret: secret,
	}
}

func (c *Config) Exchange(code string) (*Token, error) {
	rsp, err := http.Send(&http.HttpSpec{
		URL:    "https://api.weixin.qq.com/sns/oauth2/access_token",
		Method: "GET",
		URLParams: http.NewParams().Add("appid", c.AppID).Add("secret", c.Secret).
			Add("code", code).Add("grant_type", "authorization_code"),
	})

	if err != nil {
		return nil, err
	}

	var token Token
	if err = json.NewDecoder(rsp.Body).Decode(&token); err != nil {
		return nil, err
	}

	return &token, nil
}

type Token struct {
	Access_Token  string `json:"access_token"`
	Expire_In     int    `json:"expires_in"`
	Refresh_Token string `json:"refresh_token"`
	Open_ID       string `json:"openid"`
	Scope         string `json:"scope"`
}
