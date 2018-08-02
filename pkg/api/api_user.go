package api

import (
	"encoding/json"
	"errors"
	"github.com/1851616111/util/http"
	"fmt"
)

type User struct {
	Subscribe     int    `json:"subscribe"`
	OpenID        string `json:"openid"`
	NickName      string `json:"nickname"`
	Sex           int    `json:"sex"`
	Language      string `json:"language"`
	City          string `json:"city"`
	Province      string `json:"province"`
	Country       string `json:"country"`
	HeadImgUrl    string `json:"headimgurl"`
	SubscribeTime int64  `json:"subscribe_time"`
	Remark        string `json:"remark"`
	GroupID       int    `json:"groupid"`
	TagID_List    []int  `json:"tagid_list"`
}

func GetUser(token, openID string) (*User, error) {
	if len(token) == 0 || len(openID) == 0 {
		return nil, errors.New("token invalid")
	}

	rsp, err := http.Send(&http.HttpSpec{
		URL:       "https://api.weixin.qq.com/cgi-bin/user/info",
		Method:    "GET",
		URLParams: http.NewParams().Add("access_token", token).Add("openid", openID).Add("lang", "zh_CN"),
	})
	if err != nil {
		return nil, err
	}

	var res struct {
		*User
		*Error
	}
	if err := json.NewDecoder(rsp.Body).Decode(&res); err != nil {
		return nil, err
	}

	if res.Error != nil {
		return nil, errors.New(res.Error.Msg)
	} else {
		return res.User, nil
	}
}

type Error struct {
	Code int    `json:"errcode"`
	Msg  string `json:"errmsg"`
}

func (r Error) Error() error {
	if r.Code == 0 {
		return nil
	} else {
		return fmt.Errorf("request wx err: %s", r.Msg)
	}
}