package api

import (
	"github.com/1851616111/util/http"
	"fmt"
	"time"
	"github.com/pkg/errors"
	"io"
)

//临时
func CreateSceneQrCode(access_token string, expireSecs int64) (string, *QrTicket, error) {
	if expireSecs > expireSecs {
		return "", nil, errors.New("expire second too long")
	}

	var strId = fmt.Sprintf("qr_scene_%d", time.Now().UnixNano())
	var params = map[string]interface{}{
		"expire_seconds": expireSecs,
		"action_name":    "QR_STR_SCENE",
		"action_info": map[string]interface{}{
			"scene": map[string]interface{}{
				"scene_str": strId,
			},
		},
	}

	if expireSecs == 0 {
		strId = fmt.Sprintf("qr_limit_scene_%d", time.Now().UnixNano())
		params = map[string]interface{}{
			"action_name": "QR_LIMIT_STR_SCENE",
			"action_info": map[string]interface{}{
				"scene": map[string]interface{}{
					"scene_str": strId,
				},
			},
		}
	}

	spec := http.HttpSpec{
		URL:         "https://api.weixin.qq.com/cgi-bin/qrcode/create",
		Method:      "POST",
		ContentType: http.ContentType_JSON,
		URLParams:   http.NewParams().Add("access_token", access_token),
		BodyParams:  params,
	}

	var rsp QrTicket
	if err := http.NewFetcher(&spec).FetchJson(&rsp); err != nil {
		return "", nil, err
	}

	return strId, &rsp, nil
}

type QrTicket struct {
	Ticket        string `json:"ticket"`
	Url           string `json:"url"`
	ExpireSeconds int64  `json:"expire_seconds"`
}

func (t *QrTicket) GetImgReadCloser() (io.ReadCloser, error) {
	rsp, err := http.Send(&http.HttpSpec{
		URL:       "https://mp.weixin.qq.com/cgi-bin/showqrcode",
		Method:    "GET",
		URLParams: http.NewParams().Add("ticket", t.Ticket),
	})
	if err != nil {
		return nil, err
	}
	return rsp.Body, nil
}
