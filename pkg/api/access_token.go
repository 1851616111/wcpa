package api

import (
	"encoding/json"
	"errors"
	httputil "github.com/1851616111/util/http"
	"io"
)

type AccessToken struct {
	Token  string `json:"access_token,omitempty"`
	Expire uint16 `json:"expires_in,omitempty"`
	Code   int    `json:"errcode,omitempty"`
	Msg    string `json:"errmsg,omitempty"`
}

func RequestAccessToken(appId, secret string) (*AccessToken, error) {
	spec := httputil.HttpSpec{
		URL:       "https://api.weixin.qq.com/cgi-bin/token",
		Method:    "GET",
		URLParams: httputil.NewParams().Add("grant_type", "client_credential").Add("appid", appId).Add("secret", secret),
	}

	rsp, err := httputil.Send(&spec)
	if err != nil {
		return nil, err
	}

	return parseRequest(rsp.Body)
}

func parseRequest(body io.ReadCloser) (*AccessToken, error) {
	defer body.Close()

	tk := AccessToken{}
	if err := json.NewDecoder(body).Decode(&tk); err != nil {
		return nil, err
	}

	if tk.Code == 0 {
		return &tk, nil
	}

	if err := CodeErrMapping[tk.Code]; err != nil {
		return nil, err
	} else {
		return nil, errors.New(tk.Msg)
	}
}
