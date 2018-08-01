package api

import (
	"github.com/1851616111/util/http"
	"encoding/json"
	"github.com/pkg/errors"
)

type KFAccount struct {
	Id      string `json:"kf_id"`
	Account string `json:"kf_account"`
	Nick    string `json:"kf_nick"`
	HeadImg string `json:"kf_headimgurl"`
}

type NewKFAccount struct {
	Account  string `json:"kf_account"`
	NickName string `json:"nickname"`
	Password string `json:"password"`
}

func ListAllKFAccounts(access_token string) ([]KFAccount, error) {
	var res struct {
		KFList []KFAccount `json:"kf_list"`
	}
	return res.KFList, http.NewFetcher(&http.HttpSpec{
		URL:         "https://api.weixin.qq.com/cgi-bin/customservice/getkflist",
		Method:      "GET",
		ContentType: http.ContentType_JSON,
		URLParams:   http.NewParams().Add("access_token", access_token),
	}).FetchJson(&res)
}

func CreateKFAccount(access_token string, acc NewKFAccount) error {
	rsp, err := http.Send(&http.HttpSpec{
		URL:         "https://api.weixin.qq.com/customservice/kfaccount/add",
		Method:      "POST",
		ContentType: http.ContentType_JSON,
		URLParams:   http.NewParams().Add("access_token", access_token),
		BodyObject:  acc,
	})
	if err != nil {
		return err
	}

	defer rsp.Body.Close()

	var res struct {
		Code int    `json:"errcode"`
		Msg  string `json:"errmsg"`
	}
	if err := json.NewDecoder(rsp.Body).Decode(&res); err != nil {
		return err
	}

	if res.Code != 0 {
		return errors.New(res.Msg)
	}
	return nil
}
